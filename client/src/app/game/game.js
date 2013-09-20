angular.module('game', [])
    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/game/:id', {
            templateUrl:'game/game.tpl.html',
            controller:'GameCtrl'
        });
    }])
    .controller('GameCtrl', ['$scope', function($scope) {
    }]);