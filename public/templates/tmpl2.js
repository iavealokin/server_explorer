$(document).ready(function () {
    var clink = "/client"
    console.log("LINK IS")
    console.log(location.pathname)
   if(location.pathname!="/"){
    link = location.pathname
   }
    $.get(link, function (data) {
        console.log(data);
        console.log(JSON.stringify(data))
        appendResultMain(data)
    });

});
function appendResultSidebar(data) {
    var obj = jQuery.parseJSON(JSON.stringify(data));
    $("#sidebar").empty();
    $("#sidebar").append(`<p><a href="#" onclick="link('/api/getdata?path=/');">...</a></p>`)
    $.each(obj.Files, function (index, value) {
        if (value.IsFolder) {
            $("#sidebar").append(`<p><a href="#" onclick="link('/api/getdata?path=${value.Name}');">${value.Name}</a></p>`);
        }
    });
}

function appendResultMain(data) {
    var obj = jQuery.parseJSON(JSON.stringify(data));
    $("#main").empty();
    $.each(obj.Files, function (index, value) {
        if (value.IsFolder) {
          $("#main").append(`<tr><td class="left"><a href="#" onclick="link('/client${location.pathname+value.Name}');">${value.Name}</a> </td><td class="right">0</td></tr>`);
        } else {
            $("#main").append(`<tr><td class="left"><a href="/client${location.pathname+value.Name}">${value.Name}</a></td><td class="right">${value.Size}</td></tr>`);
        }
    });
}
function link(link) {
    console.log(link);
    console.log(location.host + link);
    $.get(link, function (data) {
        appendResultSidebar(data);
        appendResultMain(data);
    });
}