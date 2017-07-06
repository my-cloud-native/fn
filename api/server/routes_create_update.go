package server

import (
	"context"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab-odx.oracle.com/odx/functions/api"
	"gitlab-odx.oracle.com/odx/functions/api/models"
	"gitlab-odx.oracle.com/odx/functions/api/runner/common"
)

/* handleRouteCreateOrUpdate is used to handle POST PUT and PATCH for routes.
   Post will only create route if its not there and create app if its not.
       create only
	   Post does not skip validation of zero values
   Put will create app if its not there and if route is there update if not it will create new route.
       update if exists or create if not exists
	   Put does not skip validation of zero values
   Patch will not create app if it does not exist since the route needs to exist as well...
       update only
	   Patch accepts partial updates / skips validation of zero values.
*/
func (s *Server) handleRouteCreateOrUpdate(c *gin.Context) {
	ctx := c.MustGet("ctx").(context.Context)
	log := common.Logger(ctx)
	method := strings.ToUpper(c.Request.Method)

	var wroute models.RouteWrapper

	err, resperr := s.bindAndValidate(ctx, c, method, &wroute)
	if err != nil {
		log.WithError(err).Debug(resperr)
		c.JSON(http.StatusBadRequest, simpleError(resperr))
		return
	}

	// Create the app if it does not exist.
	err, resperr = s.ensureApp(ctx, c, wroute, method)
	if err != nil {
		log.WithError(err).Debug(resperr)
		handleErrorResponse(c, resperr)
		return
	}

	resp, err := s.updateOrInsertRoute(ctx, method, wroute)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	s.cacheRefresh(resp.Route)

	c.JSON(http.StatusOK, resp)
}

// ensureApp will only execute if it is on post or put. Patch is not allowed to create apps.
func (s *Server) ensureApp(ctx context.Context, c *gin.Context, wroute models.RouteWrapper, method string) (error, error) {
	if !(method == http.MethodPost || method == http.MethodPut) {
		return nil, nil
	}
	var app *models.App
	var err error
	app, err = s.Datastore.GetApp(ctx, wroute.Route.AppName)
	if err != nil && err != models.ErrAppsNotFound {
		return err, models.ErrAppsGet
	} else if app == nil {
		// Create a new application
		newapp := &models.App{Name: wroute.Route.AppName}
		if err = newapp.Validate(); err != nil {
			return nil, err
		}

		err = s.FireBeforeAppCreate(ctx, newapp)
		if err != nil {
			return err, models.ErrRoutesCreate
		}

		_, err = s.Datastore.InsertApp(ctx, newapp)
		if err != nil {
			return err, models.ErrRoutesCreate
		}

		err = s.FireAfterAppCreate(ctx, newapp)
		if err != nil {
			return err, models.ErrRoutesCreate
		}

	}
	return nil, nil
}

/* bindAndValidate binds the RouteWrapper to the json from the request and validates that it is correct.
If it is a put or patch it makes sure that the path in the url matches the provideed one in the body.
Defaults are set and if patch skipZero is true for validating the RouteWrapper
*/
func (s *Server) bindAndValidate(ctx context.Context, c *gin.Context, method string, wroute *models.RouteWrapper) (error, error) {
	err := c.BindJSON(wroute)
	if err != nil {
		return err, models.ErrInvalidJSON
	}

	if wroute.Route == nil {
		return err, models.ErrRoutesMissingNew
	}
	wroute.Route.AppName = c.MustGet(api.AppName).(string)

	if method == http.MethodPut || method == http.MethodPatch {
		p := path.Clean(c.MustGet(api.Path).(string))

		if wroute.Route.Path != "" && wroute.Route.Path != p {
			return nil, models.ErrRoutesPathImmutable
		}
		wroute.Route.Path = p
	}

	wroute.Route.SetDefaults()

	if err = wroute.Validate(method == http.MethodPatch); err != nil {
		return models.ErrRoutesCreate, err
	}
	return nil, nil
}

// updateOrInsertRoute will either update or insert the route respective the method.
func (s *Server) updateOrInsertRoute(ctx context.Context, method string, wroute models.RouteWrapper) (routeResponse, error) {
	var route *models.Route
	var err error
	resp := routeResponse{"Route successfully created", nil}
	up := routeResponse{"Route successfully updated", nil}

	switch method {
	case http.MethodPost:
		route, err = s.Datastore.InsertRoute(ctx, wroute.Route)
	case http.MethodPut:
		route, err = s.Datastore.UpdateRoute(ctx, wroute.Route)
		if err == models.ErrRoutesNotFound {
			// try insert then
			route, err = s.Datastore.InsertRoute(ctx, wroute.Route)
		}
	case http.MethodPatch:
		// When patching if there is an error around the app we will return one and the update fails.
		route, err = s.Datastore.UpdateRoute(ctx, wroute.Route)
		resp = up
	}
	resp.Route = route
	return resp, err
}
