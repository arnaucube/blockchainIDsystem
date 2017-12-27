'use strict';

angular.module('app.navbar', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/navbar', {
            templateUrl: 'views/navbar.html',
            controller: 'NavbarCtrl'
        });
    }])

    .controller('NavbarCtrl', function($scope, $http, $routeParams, $location) {

        $scope.user = JSON.parse(localStorage.getItem("blid_user"));

        $scope.logout = function() {
            localStorage.removeItem("blid_token");
            localStorage.removeItem("blid_user");
            window.location.reload();
        };

    });
