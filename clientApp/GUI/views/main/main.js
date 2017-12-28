'use strict';

angular.module('app.main', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/main', {
            templateUrl: 'views/main/main.html',
            controller: 'MainCtrl'
        });
    }])

    .controller('MainCtrl', function($scope, $http) {
        $scope.ids = [];
        $http.get(clientapi + 'ids')
            .then(function(data) {
                console.log('data success');
                console.log(data);
                $scope.ids = data.data;

            }, function(data) {
                console.log('data error');
            });

        $scope.newID = function() {
            $http.get(clientapi + 'newid')
                .then(function(data) {
                    console.log('data success');
                    console.log(data);
                    $scope.ids = data.data;

                }, function(data) {
                    console.log('data error');
                });
        };

        $scope.blindAndVerify = function(pubK) {
            $http.get(clientapi + 'blindandverify/' + pubK)
                .then(function(data) {
                    console.log('data success');
                    console.log(data);
                    $scope.ids = data.data;

                }, function(data) {
                    console.log('data error');
                });
        };
    });
