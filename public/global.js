'use strict';
var form_vars = {};
var alertList = [];
var alertCount = 0;
var conn;

function post_link(event)
{
	event.preventDefault();
	var form_action = $(event.target).closest('a').attr("href");
	//console.log("Form Action: " + form_action);
	$.ajax({ url: form_action, type: "POST", dataType: "json", data: {js: "1"} });
}

function bind_to_alerts() {
	$(".alertItem.withAvatar a").click(function(event) {
		event.stopPropagation();
		$.ajax({ url: "/api/?action=set&module=dismiss-alert", type: "POST", dataType: "json", data: { asid: $(this).attr("data-asid") } });
	});
}

// TODO: Add the ability for users to dismiss alerts
function load_alerts(menu_alerts)
{
	var alertListNode = menu_alerts.getElementsByClassName("alertList")[0];
	var alertCounterNode = menu_alerts.getElementsByClassName("alert_counter")[0];
	alertCounterNode.textContent = "0";
	$.ajax({
			type: 'get',
			dataType: 'json',
			url:'/api/?action=get&module=alerts',
			success: function(data) {
				if("errmsg" in data) {
					alertListNode.innerHTML = "<div class='alertItem'>"+data.errmsg+"</div>";
					return;
				}

				var alist = "";
				for(var i in data.msgs) {
					var msg = data.msgs[i];
					var mmsg = msg.msg;

					if("sub" in msg) {
						for(var i = 0; i < msg.sub.length; i++) {
							mmsg = mmsg.replace("\{"+i+"\}", msg.sub[i]);
							//console.log("Sub #" + i + ":",msg.sub[i]);
						}
					}

					if("avatar" in msg) {
						alist += "<div class='alertItem withAvatar' style='background-image:url(\""+msg.avatar+"\");'><a class='text' data-asid='"+msg.asid+"' href=\""+msg.path+"\">"+mmsg+"</a></div>";
						alertList.push("<div class='alertItem withAvatar' style='background-image:url(\""+msg.avatar+"\");'><a class='text' data-asid='"+msg.asid+"' href=\""+msg.path+"\">"+mmsg+"</a></div>");
					} else {
						alist += "<div class='alertItem'><a href=\""+msg.path+"\" class='text'>"+mmsg+"</a></div>";
						alertList.push("<div class='alertItem'><a href=\""+msg.path+"\" class='text'>"+mmsg+"</a></div>");
					}
					//console.log(msg);
					//console.log(mmsg);
				}

				if(alist == "") alist = "<div class='alertItem'>You don't have any alerts</div>";
				alertListNode.innerHTML = alist;

				if(data.msgCount != 0 && data.msgCount != undefined) {
					alertCounterNode.textContent = data.msgCount;
					menu_alerts.classList.add("has_alerts");
				} else {
					menu_alerts.classList.remove("has_alerts");
				}
				alertCount = data.msgCount;

				bind_to_alerts();
			},
			error: function(magic,theStatus,error) {
				var errtxt
				try {
					var data = JSON.parse(magic.responseText);
					if("errmsg" in data) errtxt = data.errmsg;
					else errtxt = "Unable to get the alerts";
				} catch(err) {
					errtxt = "Unable to get the alerts";
					console.log(magic.responseText);
					console.log(err);
				}
				console.log("error: ",error);
				alertListNode.innerHTML = "<div class='alertItem'>"+errtxt+"</div>";
			}
		});
}

function SplitN(data,ch,n) {
	var out = [];
	if(data.length === 0) return out;

	var lastIndex = 0;
	var j = 0;
	var lastN = 1;
	for(let i = 0; i < data.length; i++) {
		if(data[i] === ch) {
			out[j++] = data.substring(lastIndex,i);
			lastIndex = i;
			if(lastN === n) break;
			lastN++;
		}
	}
	if(data.length > lastIndex) out[out.length - 1] += data.substring(lastIndex);
	return out;
}

$(document).ready(function(){
	if(window["WebSocket"]) {
		if(window.location.protocol == "https:")
			conn = new WebSocket("wss://" + document.location.host + "/ws/");
		else conn = new WebSocket("ws://" + document.location.host + "/ws/");

		conn.onopen = function() {
			console.log("The WebSockets connection was opened");
			conn.send("page " + document.location.pathname + '\r');
			// TODO: Don't ask again, if it's denied. We could have a setting in the UCP which automatically requests this when someone flips desktop notifications on
			Notification.requestPermission();
		}
		conn.onclose = function() {
			conn = false;
			console.log("The WebSockets connection was closed");
		}
		conn.onmessage = function(event) {
			//console.log("WS_Message: ",event.data);
			if(event.data[0] == "{") {
				try {
					var data = JSON.parse(event.data);
				} catch(err) {
					console.log(err);
				}

				if ("msg" in data) {
					var msg = data.msg
					if("sub" in data)
						for(var i = 0; i < data.sub.length; i++)
							msg = msg.replace("\{"+i+"\}", data.sub[i]);

					if("avatar" in data) alertList.push("<div class='alertItem withAvatar' style='background-image:url(\""+data.avatar+"\");'><a class='text' data-asid='"+data.asid+"' href=\""+data.path+"\">"+msg+"</a></div>");
					else alertList.push("<div class='alertItem'><a href=\""+data.path+"\" class='text'>"+msg+"</a></div>");
					if(alertList.length > 8) alertList.shift();
					//console.log("post alertList",alertList);
					alertCount++;

					var alist = ""
					for (var i = 0; i < alertList.length; i++) alist += alertList[i];

					//console.log(alist);
					// TODO: Add support for other alert feeds like PM Alerts
					var general_alerts = document.getElementById("general_alerts");
					var alertListNode = general_alerts.getElementsByClassName("alertList")[0];
					var alertCounterNode = general_alerts.getElementsByClassName("alert_counter")[0];
					alertListNode.innerHTML = alist;
					alertCounterNode.textContent = alertCount;

					// TODO: Add some sort of notification queue to avoid flooding the end-user with notices?
					// TODO: Use the site name instead of "Something Happened"
					if(Notification.permission === "granted") {
						var n = new Notification("Something Happened",{
							body: msg,
							icon: data.avatar,
						});
						setTimeout(n.close.bind(n), 8000);
					}

					bind_to_alerts();
				}
			}

			var messages = event.data.split('\r');
			for(var i = 0; i < messages.length; i++) {
				//console.log("Message: ",messages[i]);
				if(messages[i].startsWith("set ")) {
					//msgblocks = messages[i].split(' ',3);
					let msgblocks = SplitN(messages[i]," ",3);
					if(msgblocks.length < 3) continue;
					document.querySelector(msgblocks[1]).innerHTML = msgblocks[2];
				} else if(messages[i].startsWith("set-class ")) {
					let msgblocks = SplitN(messages[i]," ",3);
					if(msgblocks.length < 3) continue;
					document.querySelector(msgblocks[1]).className = msgblocks[2];
				}
			}
		}
	}
	else conn = false;

	$(".open_edit").click(function(event){
		//console.log("clicked on .open_edit");
		event.preventDefault();
		$(".hide_on_edit").hide();
		$(".show_on_edit").show();
	});

	$(".topic_item .submit_edit").click(function(event){
		event.preventDefault();
		//console.log("clicked on .topic_item .submit_edit");
		$(".topic_name").html($(".topic_name_input").val());
		$(".topic_content").html($(".topic_content_input").val());
		$(".topic_status_e:not(.open_edit)").html($(".topic_status_input").val());

		$(".hide_on_edit").show();
		$(".show_on_edit").hide();

		let topicNameInput = $('.topic_name_input').val();
		let topicStatusInput = $('.topic_status_input').val();
		let topicContentInput = $('.topic_content_input').val();
		let formAction = this.form.getAttribute("action");
		//console.log("New Topic Name: ", topicNameInput);
		//console.log("New Topic Status: ", topicStatusInput);
		//console.log("New Topic Content: ", topicContentInput);
		//console.log("Form Action: ", formAction);
		$.ajax({
			url: formAction,
			type: "POST",
			dataType: "json",
			data: {
				topic_name: topicNameInput,
				topic_status: topicStatusInput,
				topic_content: topicContentInput,
				topic_js: 1
			}
		});
	});

	$(".delete_item").click(function(event)
	{
		post_link(event);
		$(this).closest('.deletable_block').remove();
	});

	$(".edit_item").click(function(event)
	{
		event.preventDefault();
		let blockParent = $(this).closest('.editable_parent');
		let block = blockParent.find('.editable_block').eq(0);
		block.html("<textarea style='width: 99%;' name='edit_item'>" + block.html() + "</textarea><br /><a href='" + $(this).closest('a').attr("href") + "'><button class='submit_edit' type='submit'>Update</button></a>");

		$(".submit_edit").click(function(event)
		{
			event.preventDefault();
			let blockParent = $(this).closest('.editable_parent');
			let block = blockParent.find('.editable_block').eq(0);
			let newContent = block.find('textarea').eq(0).val();
			block.html(newContent);

			var formAction = $(this).closest('a').attr("href");
			//console.log("Form Action:",formAction);
			$.ajax({ url: formAction, type: "POST", dataType: "json", data: { isJs: "1", edit_item: newContent }
			});
		});
	});

	$(".edit_field").click(function(event)
	{
		event.preventDefault();
		let blockParent = $(this).closest('.editable_parent');
		let block = blockParent.find('.editable_block').eq(0);
		block.html("<input name='edit_field' value='" + block.text() + "' type='text'/><a href='" + $(this).closest('a').attr("href") + "'><button class='submit_edit' type='submit'>Update</button></a>");

		$(".submit_edit").click(function(event) {
			event.preventDefault();
			let blockParent = $(this).closest('.editable_parent');
			let block = blockParent.find('.editable_block').eq(0);
			let newContent = block.find('input').eq(0).val();
			block.html(newContent);

			let formAction = $(this).closest('a').attr("href");
			//console.log("Form Action:", formAction);
			$.ajax({
				url: formAction + "?session=" + session,
				type: "POST",
				dataType: "json",
				data: { isJs: "1", edit_item: newContent }
			});
		});
	});

	$(".edit_fields").click(function(event)
	{
		event.preventDefault();
		if($(this).find("input").length !== 0) return;
		//console.log("clicked .edit_fields");
		var block_parent = $(this).closest('.editable_parent');
		//console.log(block_parent);
		block_parent.find('.hide_on_edit').hide();
		block_parent.find('.show_on_edit').show();
		block_parent.find('.editable_block').show();
		block_parent.find('.editable_block').each(function(){
			var field_name = this.getAttribute("data-field");
			var field_type = this.getAttribute("data-type");
			if(field_type=="list")
			{
				var field_value = this.getAttribute("data-value");
				if(field_name in form_vars) var it = form_vars[field_name];
				else var it = ['No','Yes'];
				var itLen = it.length;
				var out = "";
				//console.log("Field Name:",field_name);
				//console.log("Field Type:",field_type);
				//console.log("Field Value:",field_value);
				for (var i = 0; i < itLen; i++) {
					var sel = "";
					if(field_value == i || field_value == it[i]) {
						sel = "selected ";
						this.classList.remove(field_name + '_' + it[i]);
						this.innerHTML = "";
					}
					out += "<option "+sel+"value='"+i+"'>"+it[i]+"</option>";
				}
				this.innerHTML = "<select data-field='"+field_name+"' name='"+field_name+"'>"+out+"</select>";
			}
			else if(field_type=="hidden") {}
			else this.innerHTML = "<input name='"+field_name+"' value='"+this.textContent+"' type='text'/>";
		});

		// Remove any handlers already attached to the submitter
		$(".submit_edit").unbind("click");

		$(".submit_edit").click(function(event)
		{
			event.preventDefault();
			//console.log("running .submit_edit event");
			var out_data = {isJs: "1"}
			var block_parent = $(this).closest('.editable_parent');
			block_parent.find('.editable_block').each(function() {
				var field_name = this.getAttribute("data-field");
				var field_type = this.getAttribute("data-type");
				if(field_type=="list") {
					var newContent = $(this).find('select :selected').text();
					this.classList.add(field_name + '_' + newContent);
					this.innerHTML = "";
				} else if(field_type=="hidden") {
					var newContent = $(this).val();
				} else {
					var newContent = $(this).find('input').eq(0).val();
					this.innerHTML = newContent;
				}
				this.setAttribute("data-value",newContent);
				out_data[field_name] = newContent;
			});

			var form_action = $(this).closest('a').attr("href");
			//console.log("Form Action:", form_action);
			//console.log(out_data);
			$.ajax({ url: form_action + "?session=" + session, type:"POST", dataType:"json", data: out_data });
			block_parent.find('.hide_on_edit').show();
			block_parent.find('.show_on_edit').hide();
		});
	});

	// This one's for Tempra Conflux
	// TODO: We might want to use pure JS here
	$(".ip_item").each(function(){
		var ip = this.textContent;
		if(ip.length > 10){
			this.innerHTML = "Show IP";
			this.onclick = function(event){
				event.preventDefault();
				this.textContent = ip;
			};
		}
	});

	$(this).click(function() {
		$(".selectedAlert").removeClass("selectedAlert");
		$("#back").removeClass("alertActive");
	});
	$(".alert_bell").click(function(){
		var menu_alerts = $(this).parent();
		if(menu_alerts.hasClass("selectedAlert")) {
			event.stopPropagation();
			menu_alerts.removeClass("selectedAlert");
			$("#back").removeClass("alertActive");
		}
	});

	var alert_menu_list = document.getElementsByClassName("menu_alerts");
	for(var i = 0; i < alert_menu_list.length; i++) {
		load_alerts(alert_menu_list[i]);
	}

	$(".menu_alerts").click(function(event) {
		event.stopPropagation();
		if($(this).hasClass("selectedAlert")) return;
		if(!conn) load_alerts(this);
		this.className += " selectedAlert";
		document.getElementById("back").className += " alertActive"
	});

	$("input,textarea,select,option").keyup(function(event){
		event.stopPropagation();
	})

	$(".create_topic_link").click(function(event){
		event.preventDefault();
		$(".topic_create_form").show();
	});
	$(".topic_create_form .close_form").click(function(){
		event.preventDefault();
		$(".topic_create_form").hide();
	});

	function uploadFileHandler() {
		var fileList = this.files;

		// Truncate the number of files to 5
		let files = [];
		for(var i = 0; i < fileList.length && i < 5; i++)
			files[i] = fileList[i];

		// Iterate over the files
		for(let i = 0; i < files.length; i++) {
			console.log("files[" + i + "]",files[i]);
			let reader = new FileReader();
			reader.onload = function(e) {
				var fileDock = document.getElementById("upload_file_dock");
				var fileItem = document.createElement("label");
				console.log("fileItem",fileItem);

				if(!files[i]["name"].indexOf('.' > -1)) {
					// TODO: Surely, there's a prettier and more elegant way of doing this?
					alert("This file doesn't have an extension");
					return;
				}

				var ext = files[i]["name"].split('.').pop();
				fileItem.innerText = "." + ext;
				fileItem.className = "formbutton uploadItem";
				fileItem.style.backgroundImage = "url("+e.target.result+")";

				fileDock.appendChild(fileItem);

				let reader = new FileReader();
				reader.onload = function(e) {
					crypto.subtle.digest('SHA-256',e.target.result).then(function(hash) {
						const hashArray = Array.from(new Uint8Array(hash))
						return hashArray.map(b => ('00' + b.toString(16)).slice(-2)).join('')
					}).then(function(hash) {
						console.log("hash",hash);
						let content = document.getElementById("input_content")
						console.log("content.value",content.value);
						
						if(content.value == "") content.value = content.value + "//" + siteURL + "/attachs/" + hash + "." + ext;
						else content.value = content.value + "\r\n//" + siteURL + "/attachs/" + hash + "." + ext;
						console.log("content.value",content.value);
					});
				}
				reader.readAsArrayBuffer(files[i]);
			}
			reader.readAsDataURL(files[i]);
		}
	}

	var uploadFiles = document.getElementById("upload_files");
	if(uploadFiles != null) {
		uploadFiles.addEventListener("change", uploadFileHandler, false);
	}

	$("#themeSelectorSelect").change(function(){
		console.log("Changing the theme to " + this.options[this.selectedIndex].getAttribute("val"));
		$.ajax({
			url: this.form.getAttribute("action") + "?session=" + session,
			type: "POST",
			dataType: "json",
			data: { "newTheme": this.options[this.selectedIndex].getAttribute("val"), isJs: "1" },
			success: function (data, status, xhr) {
				console.log("Theme successfully switched");
				console.log("data",data);
				console.log("status",status);
				console.log("xhr",xhr);
				window.location.reload();
			},
			// TODO: Use a standard error handler for the AJAX calls in here which throws up the response (if JSON) in a .notice? Might be difficult to trace errors in the console, if we reuse the same function every-time
			error: function(xhr,status,errstr) {
				console.log("The AJAX request failed");
				console.log("xhr",xhr);
				console.log("status",status);
				console.log("errstr",errstr);
				if(status=="parsererror") {
					console.log("The server didn't respond with a valid JSON response");
				}
			}
		});
	});

	this.onkeyup = function(event) {
		if(event.which == 37) this.querySelectorAll("#prevFloat a")[0].click();
		if(event.which == 39) this.querySelectorAll("#nextFloat a")[0].click();
	};
});
