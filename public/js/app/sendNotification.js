messageElement = $("#message")
tokenElement = $("#token")
notificationElement = $("#notification")
errorElement = $('#error');
firebase.initializeApp({
    'messagingSenderId': 'BN5cHhpN8vO_bWKQ9LwvT5vaMTep6W1O4DOoA-DCjwDSkreGTFzKBKCFM3TAtNGAJ1oAdsTdrxguxG_HhiVVu9I'
});
const messaging = firebase.messaging();

function initFirebaseMessagingRegistration() {
    messaging
        .requestPermission()
        .then(function () {
            messageElement.innerHTML = "Got notification permission";
            console.log("Got notification permission");
            return messaging.getToken();
        })
        .then(function (token) {
            // print the token on the HTML page
            tokenElement.innerHTML = "Token is " + token;
        })
        .catch(function (err) {
            errorElement.innerHTML = "Error: " + err;
            console.log("Didn't get notification permission", err);
        });
}
messaging.onMessage(function (payload) {
    console.log("Message received. ", JSON.stringify(payload));
    notificationElement.innerHTML = notificationElement.innerHTML + " " + payload.data.notification;
});
messaging.onTokenRefresh(function () {
    messaging.getToken()
        .then(function (refreshedToken) {
            console.log('Token refreshed.');
            tokenElement.innerHTML = "Token is " + refreshedToken;
        }).catch(function (err) {
        errorElement.innerHTML = "Error: " + err;
        console.log('Unable to retrieve refreshed token ', err);
    });
});