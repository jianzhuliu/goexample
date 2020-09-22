function lineHighlight(error) {
	var regex = /prog.go:([0-9]+)/g;
	var r = regex.exec(error);
	while (r) {
		$(".lines div").eq(r[1] - 1).addClass("lineerror");
		r = regex.exec(error);
	}
}


function lineClear() {
	$(".lineerror").removeClass("lineerror");
}

function loading() {
	lineClear();
	$("#output").removeClass("error").text('Waiting for remote server...');
}

function body() {
	return $("#code").val();
}

function setBody(text) {
	$("#code").val(text);
}

function setError(error) {
	lineClear();
	lineHighlight(error);
	$("#output").empty().addClass("error").text(error);
}

function setOutput(text){
	lineClear();
	$("#output").empty().removeClass("error").text(text);
}

function run() {
	loading();
	var data = {
		"body": body()
	};

	$.ajax("/run", {
		data: data,
		type: "POST",
		dataType: "json",
		success: function(data) {
			if (data.Error) {
				setError(data.Error);
			} else {
				setOutput(data.Body);
			}
		}
	});
}

function fmt() {
	loading();
	var data = {
		"body": body()
	};

	$.ajax("/fmt", {
		data: data,
		type: "POST",
		dataType: "json",
		success: handleData
	});
}

function handleData(data){
	if (data.Error) {
		setError(data.Error);
	} else {
		setBody(data.Body);
		setError("");
	}
	
}

function imports() {
	loading();
	var data = {
		"body": body()
	};

	$.ajax("/imports", {
		data: data,
		type: "POST",
		dataType: "json",
		success: handleData
	});
}