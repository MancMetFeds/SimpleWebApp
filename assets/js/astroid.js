var ship;

function begin(){
    astroids.start()
    ship = new sprite(document.getElementById("area").offsetWidth/2, document.getElementById("area").offsetHeight/2, "ship", "white")
    console.log(document.getElementById("area").offsetWidth/2, document.getElementById("area").offsetHeight/2);
    
}

var astroids = {
    canvas : document.createElement("canvas"),
    start: function() {
        var introText = "What the-? We've come out of the Internet, right into a meteor shower... or an asteroid field or something. It's not on any of the charts!"
        this.canvas.id ="nav-canvas";
        this.canvas.width = document.getElementById("area").offsetWidth;
        this.canvas.height = document.getElementById("area").offsetHeight;
        this.context = this.canvas.getContext("2d");
        this.context.font = "30px Arial";
        this.context.strokeStyle='white';
        this.context.strokeText(introText, (this.canvas.width-this.context.measureText(introText).width)/2, 100); 
        document.body.appendChild(this.canvas);
        
    }    
}    

function sprite(startx,starty, spriteName, colour){
    var canvas = document.getElementById("nav-canvas");
    var context = canvas.getContext("2d");
    context.beginPath();
    context.moveTo(startx, starty)  // firstpoint
    context.fillStyle = colour;
    switch (spriteName){
        case "ship":
            context.lineTo(startx+20, starty+50);   //left
            context.lineTo(startx, starty+30);      //bottom-center
            context.lineTo(startx-20, starty+50);   //right
            
        //add other game entities here
          
    
    }

    context.fill();

    function movement(event){
        console.log(event)
        var key_Code=event.keyCode;
        if (this.spriteName = "ship"){
            switch (key_Code){
                case 38:
                    this.startx += this.startx-5;
                    console.log(this.startx)
                case 39:
                case 40:
                case 37:
            }
        }
    }
}




