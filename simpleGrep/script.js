'use strict';

function loadTextArea(data) {
    var textArea = document.getElementById("text_box");
    textArea.value = data;
}

function setPath(path_selected) {
    var path = document.getElementById("path");
    path.value = path_selected;
}

function updateStatus(data) {
    var statusLbl = document.getElementById("statusLbl");
    statusLbl.innerHTML = data;
}

function callSearch() {
    var pattern = document.getElementById("pattern");
    var path = document.getElementById("path");
    var caseInsensitive = document.getElementById("caseInsensitive");
    var wholeWord = document.getElementById("wholeWord");					
    var wholeLine = document.getElementById("wholeLine");
    var filenameOnly = document.getElementById("filenameOnly");
    var inverted = document.getElementById("inverted");
    var options = {caseInsensitive:caseInsensitive.checked, wholeWord:wholeWord.checked, wholeLine:wholeLine.checked, filenameOnly:filenameOnly.checked, inverted:inverted.checked};
    if (pattern.value.trim().length === 0) {
        error("No pattern entered!");
        return;
    }
    if (path.value.trim().length === 0) {
        error("No path entered!");
        return;
    }
    search(pattern.value, path.value, options)
}
