
const boxes = 10;
const colors = ["red", "orange", "yellow", "green", "blue", "indigo", "violet"]


function reset() {
    $(".bar-cell").addClass("bar-cell-active");	
    }

function activateBoxes(percentage) {
    var value = Math.ceil(parseFloat(percentage) * boxes);
    for (var i = 0; i < value; i++) {
        $(".bar-cell").eq(i).addClass("bar-cell-active");    
    }
    for (var i = value; i < boxes; i++) {
        $(".bar-cell").eq(i).removeClass("bar-cell-active");    
    }
}

function colorBox(index, color) {
    $(".bar-cell").eq(index).css('background-color', color);
}

function colorBoxes(segments) {
    var value = Math.ceil(parseFloat(percentage) * boxes);
    for (var i = 0; i < value; i++) {
        $(".bar-cell").eq(i).addClass("bar-cell-active");    
    }
}

function updateProgressBar() {
    $.get(
        "/api/button",
        {},
        function(data) {
            data.Segments.forEach(function(element) {
                //Color the amount of boxes, with a different color for each segment
            })
            activateBoxes(data.Percentage)
        },
        "json"
    )
}

$(document).ready(function() {
    $(".bar-cell").remove();
    for (let i = 0; i < boxes; i++) {
        $(".bar").append("<div class='bar-cell'></div>");
    }
    updateProgressBar()
    setInterval(updateProgressBar, 5000)

    $(".game-button").click(function() {
        $.post(
            "/api/button",
            {},
            function(data) {
                alert("Clicked at: " + data.percentage);
                console.log(data);
                reset()
            },
            "json"
         )
    });
})
