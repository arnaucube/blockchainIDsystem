'use strict';

angular.module('app.login', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/login', {
            templateUrl: 'views/login/login.html',
            controller: 'LoginCtrl'
        });
    }])

    .controller('LoginCtrl', function($scope, $http, $routeParams, toastr) {
        $scope.user = {};
        $scope.login = function() {
            console.log('Doing login', $scope.user);
            console.log(urlapi + "login");
            $http({
                    url: urlapi + 'login',
                    method: "POST",
                    headers: {
                        "Content-Type": undefined
                    },
                    data: $scope.user
                })
                .then(function(data) {
                        console.log("data: ");
                        console.log(data.data);
                        if (data.data.token) {
                            localStorage.setItem("blid_token", data.data.token);
                            localStorage.setItem("blid_user", JSON.stringify(data.data));
                            window.location.reload();
                        } else {
                            console.log("login failed");
                            toastr.error('Login failed');
                        }


                    },
                    function(data) {
                        console.log(data);
                    });

        };
    });
