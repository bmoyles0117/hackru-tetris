
tetrisApp = angular.module("tetrisApp" , [])
  // .config(function($interpolateProvider) {
  // 	  $interpolateProvider.startSymbol('[[');
  // 	  $interpolateProvider.endSymbol(']]');
  // })
tetrisApp.directive('gameStarter', function() {
  return function($scope, element, attrs) {
    console.log(element);

    element.bind('click', function() {
      $.post('/start');
      
      return false;
    });
  }
})
tetrisApp.controller("HomepageCtrl",["$scope", function($scope) {
    $scope.game_over = true;
    $scope.game_over1 = true;
    $scope.game_over2 = true;

    $scope.loading = true;
    $scope.teams = 1;


    $scope.colors = {
          73 : "blue",
          74 : "green",
          76 : "yellow",
          79 : "red",
          83 : "cyan",
          84 : "black",
          90 : "gray"
       };


    $scope.pusher = new Pusher('c2388f10a4afc865f3a5');
    $scope.channel = $scope.pusher.subscribe('my_channel');


    $scope.channel.bind('my_event', function(data) {
    	$scope.gameData = data.cells;
      $scope.game_over = data.game_over;
      // $scope.gameNextPiece = data.next_piece;

      if(data.game_over) {
        alert('GAME OVER');
        return
      }


      $scope.drawPieces = function () {

        if( data.board_key == "+17327305402"  ){
          var canvas = document.getElementById('myCanvas');
        }else {
          var canvas = document.getElementById('myCanvas3');
        }

        var context = canvas.getContext('2d');

        canvas.width = canvas.width;

        for(i=0; i<22; i++){
        	for(j=0; j<12; j++ ) {

            if ($scope.gameData[i][j] != 0){
    					context.fillStyle=$scope.colors[$scope.gameData[i][j]];
              context.fillRect( ((canvas.width/12)*j) ,((canvas.height/22)*i),canvas.width/12, canvas.height/22);
               context.strokeStyle="#f9f9f9";
              context.strokeRect( ((canvas.width/12)*j) ,((canvas.height/22)*i),canvas.width/12, canvas.height/22);

    				}





        	}

        }
        if(data.next_piece){
          piece_size = (data.next_piece.shape[0]).length;
        
        if( data.board_key == "+17327305402"  ){
          var canvas = document.getElementById('myCanvas2');
        }else {
          var canvas = document.getElementById('myCanvas4');
        }
          var context = canvas.getContext('2d');

        canvas.width = canvas.width;

        console.log(piece_size);

        for(i=0; i<piece_size; i++){
          for(j=0; j<piece_size; j++ ) {

            if (data.next_piece.shape[i][j] != 0){
              context.fillStyle=$scope.colors[data.next_piece.color];

              context.fillRect( ((canvas.width/piece_size)*j) ,((canvas.height/piece_size)*i),canvas.width/piece_size, canvas.height/piece_size);
              context.strokeStyle="#f9f9f9";
              context.strokeRect( ((canvas.width/piece_size)*j) ,((canvas.height/piece_size)*i),canvas.width/piece_size, canvas.height/piece_size);
 
            }
        }


	     }
    }

    }

    	$scope.drawPieces(data.board_key);



    });



}]);



  // $scope.drawPieces = function () {
  //       var canvas = document.getElementById('myCanvas');
  //       var context = canvas.getContext('2d');

 

  // }

