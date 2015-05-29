
angular.module("FF", []).controller("HelloCtrl",
	function($scope){
	$scope.cname = "ANGULARJS_F";
	$scope.gName = function () {
		return $scope.cname+"_gName";
	}

	$scope.kids = [
		{name:"Nacy",age:25	},
		{name:"Oari",age:27	},
		{name:"Oari",age:27	},
		{name:"Oari",age:27	},
		{name:"Oari",age:27	},
		{name:"Oari",age:27	},
		{name:"Oari",age:27	},
		{name:"Oari",age:27	},
		{name:"Oari",age:27	},
	];
});
angular.module("app", []).controller("HelloCtrl",
	function ($scope) {
    $scope.cname="ANGULARJS_A";
});

