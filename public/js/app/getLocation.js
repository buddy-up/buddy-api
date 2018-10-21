var x = $('#demo');

function getLocation() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(showPosition);
    } else {
        x.html("aaa");
    }
}
function showPosition(position) {
    var x = $('#demo');
    x.html("Latitude: " + position.coords.latitude +
        "<br>Longitude: " + position.coords.longitude);
}