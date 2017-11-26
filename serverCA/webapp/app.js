'use strict';

var urlapi = "http://localhost:3030/";

// Declare app level module which depends on views, and components
angular.module('webapp', [
  'ngRoute',
  'ngMessages',
  'angularBootstrapMaterial',
  'toastr',
  'app.navbar',
  'app.main'
]).
config(['$locationProvider', '$routeProvider', function($locationProvider, $routeProvider) {
  $locationProvider.hashPrefix('!');

  $routeProvider.otherwise({redirectTo: '/main'});
}])
.config(function(toastrConfig) {
  angular.extend(toastrConfig, {
    autoDismiss: false,
    containerId: 'toast-container',
    maxOpened: 0,
    newestOnTop: true,
    positionClass: 'toast-bottom-right',
    preventDuplicates: false,
    preventOpenDuplicates: false,
    target: 'body'
  });
})
.factory('httpInterceptor', function httpInterceptor () {
  return {
    request: function(config) {
      return config;
    },

    requestError: function(config) {
      return config;
    },

    response: function(res) {
      return res;
    },

    responseError: function(res) {
      return res;
    }
  };
})
.factory('api', function ($http) {
	return {
		init: function () {
      /*$http.defaults.headers.common['X-Access-Token'] = localStorage.getItem('cr_webapp_token');
      $http.defaults.headers.post['X-Access-Token'] = localStorage.getItem('cr_webapp_token');*/
		}
	};
})
.run(function (api) {
	api.init();
});
