angular.module('game', [])
    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/game/:id', {
            templateUrl:'game/game.tpl.html',
            controller:'GameCtrl'
        });
    }])
    .controller('GameCtrl', ['$scope', function($scope) {
        $scope.game = {
            Teams: [
                {
                    Players: [
                        {
                            Name: 'Reed',
                            Id: 'abcd-abcd-abcd-abcd',
                            Hand: [0,1,2,3,4]
                        },
                        {
                            Name: 'Adina',
                            Id: 'abcd-abcd-abcd-abcd',
                            Hand: [-1, -1, -1, -1, -1]
                        }
                    ],
                    Score: 4,
                    HandsWon: 2
                },
                {
                    Players: [
                        {
                            Name: 'Joe',
                            Id: 'abcd-abcd-abcd-abcd',
                            Hand: [-1, -1, -1, -1, -1]
                        },
                        {
                            Name: 'Serena',
                            Id: 'abcd-abcd-abcd-abcd',
                            Hand: [-1, -1, -1, -1, -1]
                        }
                    ],
                    Score: 3,
                    HandsWon: 2
                }
            ],
            Trump: 0
        };

    }]);