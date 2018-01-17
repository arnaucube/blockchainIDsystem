'use strict';

angular.module('app.navbar', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/navbar', {
            templateUrl: 'views/navbar.html',
            controller: 'NavbarCtrl'
        });
    }])

    .controller('NavbarCtrl', function($scope, $rootScope, $http, $routeParams, $location) {
        $rootScope.server = JSON.parse(localStorage.getItem("old_darkID_server"));

        $scope.user = JSON.parse(localStorage.getItem("old_darkID_user"));

        $scope.logout = function() {
            localStorage.removeItem("old_darkID_token");
            localStorage.removeItem("old_darkID_user");
            localStorage.removeItem("old_darkID_server");
            window.location.reload();
        };

    });
