$(document).ready(function(){
	// TODO: Run this when the image is loaded rather than when the document is ready?
	$(".topic_list img").each(function(){
		let aspectRatio = this.naturalHeight / this.naturalWidth;
		console.log("aspectRatio ",aspectRatio);
		console.log("this.height ",this.naturalHeight);
		console.log("this.width ",this.naturalWidth);

		$(this).css({
			height: aspectRatio * this.width
		});
	});
});