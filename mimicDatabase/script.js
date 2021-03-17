"use strict";

function callConnect() {
  let password = document.getElementById("password").value;
  connect(password);
  document.getElementById("callConnectBtn").disabled = true;
  document.getElementById("callQueryBtn").disabled = false;
  document.getElementById("password").disabled = true;
}

function updateTextArea(data) {
  document.getElementById("textarea").value = data;
}

function callQuery() {
  let matchType = getMatchType();
  let queryValue = document.getElementById("queryValue").value;
  query(matchType, queryValue);
}

function getMatchType() {
  if (document.getElementById("exactMatch").checked) {
    return "EXACT";
  } else if (document.getElementById("likeMatch").checked) {
    return "LIKE";
  }
}