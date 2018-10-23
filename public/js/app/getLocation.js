var x = $('#demo');

function getLocation() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(showPosition);
        navigator.geolocation.getCurrentPosition(saveLocation);
    } else {
        x.html("aaa");
    }
}
function showPosition(position) {
    var x = $('#demo');
    x.html("Latitude: " + position.coords.latitude +
        "<br>Longitude: " + position.coords.longitude);
}

function saveLocation(position){
    $.ajax({
        type: "POST",
        url: "/checkin",
        data:{
            "latitude" : position.coords.latitude,
            "longitude": position.coords.longitude,
        }
    });
}