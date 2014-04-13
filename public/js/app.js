
tetrisApp = angular.module("tetrisApp" , [])
  // .config(function($interpolateProvider) {
  //   $interpolateProvider.startSymbol('[[');
  //   $interpolateProvider.endSymbol(']]');
  // })
tetrisApp.directive('gameStarter', function() {
  return function($scope, element, attrs) {
    element.bind('click', function() {
      $.post('/start');
      
      return false;
    });
  }
});
tetrisApp.controller("HomepageCtrl",["$scope", function($scope) {
  $scope.game_over = true;
  $scope.loading = true;
  $scope.teams = 2;

  $scope.colors = {
    73 : "#77dcba",
    74 : "#456cf7",
    76 : "#1931a8",
    79 : "#f59c19",
    83 : "#d7beb8",
    84 : "#f970e7",
    90 : "#da3951"
  };

  var pusher = new Pusher('c2388f10a4afc865f3a5');
  var channel = pusher.subscribe('my_channel');

  channel.bind('my_event', function(data) {
    $scope.game_over = data.game_over;

    if(data.game_over) {
      return
    }

    var canvas = document.getElementById(data.board_key == "+17327305402" ? 'player1_board' : 'player2_board');
    var context = canvas.getContext('2d');

    canvas.width = canvas.width;

    for(var i = 0; i < data.cells.length; i++){
      for(var j = 0; j < data.cells[i].length; j++ ) {
        var rows = data.cells.length,
            cols = data.cells[i].length,
            width = canvas.width/cols,
            height = canvas.height/rows;

        if (data.cells[i][j] != 0) {
          context.fillStyle = $scope.colors[data.cells[i][j]];
          context.fillRect(width*j, height*i, width, height);
          context.strokeStyle = "#f9f9f9";
          context.strokeRect(width*j, height*i, width, height);
        }
      }
    }

    if(data.next_piece){
      var piece_size = (data.next_piece.shape[0]).length;
      var canvas = document.getElementById(data.board_key == "+17327305402" ? 'player1_next' : 'player2_next');
      var context = canvas.getContext('2d');

      canvas.width = canvas.width;

      for(var i = 0; i < piece_size; i++){
        for(var j = 0; j < piece_size; j++ ) {
          if (data.next_piece.shape[i][j] != 0){
            var width = canvas.width/piece_size,
                height = canvas.height/piece_size;

            context.fillStyle = $scope.colors[data.next_piece.color];
            context.fillRect(width*j, height*i, width, height);
            context.strokeStyle="#f9f9f9";
            context.strokeRect(width*j, height*i, width, height);
          }
        }
      }
    }
  });
}]);