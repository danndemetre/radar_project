$(document).read( function() {

	$.getJSON("fetch.php", function(data) {
		$("ul").empty();

		$.each(data.result, function(){
			$("ul").append("<li>Name: "+this['name']+"</li><li>Age: "+this['age']+"<li>Company: "+this['company']+"</li><li><br />";
		});

    });

});
