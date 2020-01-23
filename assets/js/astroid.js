var ship

function begin(){
    astroids.start()
    ship = new ships;
    ship.place(document.getElementById("area").offsetWidth/2, document.getElementById("area").offsetHeight/2, "ship", "white");        
    document.onkeydown = (e) => {ship.accelerate(e);};
    document.onkeyup = (e) => {ship.deaccelerate(e);};
    console.log(document.getElementById("area").offsetWidth/2, document.getElementById("area").offsetHeight/2);
    
}




var astroids = {
    canvas : document.createElement("canvas"),
    start: function() {
        var introText = "What the-? We've come out of the Internet, right into a meteor shower... or an asteroid field or something. It's not on any of the charts!"
        this.canvas.id ="nav-canvas";
        this.canvas.width = document.getElementById("area").offsetWidth;

        this.canvas.height = document.getElementById("area").offsetHeight;
        this.canvas.style.width = "100%"; // Note you must post fix the unit type %,px,em
        this.canvas.style.height = "100%";
        this.context = this.canvas.getContext("2d");
        this.context.font = "30px Arial";
        this.context.strokeStyle='white';
        this.context.strokeText(introText, (this.canvas.width-this.context.measureText(introText).width)/2, 100); 
        document.body.appendChild(this.canvas);
        
    }
    
}    

function ships() {
    this.x = null,
    this.y = null,
    this.name = null,
    this.colour = null,
    this.canvas = document.getElementById("nav-canvas"),
    this.context = this.canvas.getContext("2d");
    this.minSpeed=0.5
    this.maxSpeed=20
    this.acceleration=0.02
    this.currentspeed=0
    this.deacceleration=0.05
    //place object on canvas
    this.newparams = function (pointx, pointy) {
        
        this.x = pointx;
        this.y = pointy;
        console.log(this.x, this.y, "...",pointx, pointy)

    };

    this.place = function (startx, starty, name, colour) {
        this.name = name;
        this.colour = colour;
        this.newparams(startx, starty)
        
        this.context.beginPath();
        this.context.moveTo(this.x, this.y);  // firstpoint
        this.context.fillStyle = colour;
        this.context.lineTo(this.x+20, this.y+50);   //left
        this.context.lineTo(this.x, this.y+30);      //bottom-center
        this.context.lineTo(this.x-20, this.y+50);   //right
        this.context.fill()
    };

    //clear object from canvas
    this.clear = function () {
        this.context.clearRect(0, 0, this.canvas.width, this.canvas.height);
    };
    
    //update object position on canvas
    this.update = function (){
        this.clear();
        this.context.beginPath();
        this.context.moveTo(this.x, this.y);  // firstpoint
        this.context.fillStyle = this.colour;
        this.context.lineTo(this.x+20, this.y+50);   //left
        this.context.lineTo(this.x, this.y+30);      //bottom-center
        this.context.lineTo(this.x-20, this.y+50);   //right
        this.context.fill()
    };

    this.movement = function(event, distance){
        

        if (event.key == "ArrowUp"){
            this.newparams(this.x, (this.y-=distance));
        } if (event.key == "ArrowRight"){
            this.newparams((this.x+=distance), this.y);
        } if (event.key == "ArrowLeft"){
            this.newparams((this.x-=distance), this.y);
        } if (event.key == "ArrowDown"){
            this.newparams(this.x, (this.y+=distance));
        }
        this.update()
    };

    this.accelerate = function(event){
        if (this.currentspeed == 0 ){
            this.currentspeed = 0.5;
            this.movement(event, currentspeed);
            return;
        }
        this.speed*=this.acceleration;
        this.movement(event, currentspeed);
            return;
        

    }

    this.deaccelerate = function(event){
        setInterval(function(){ console.log(this.name)}, 3000);
        // if (this.currentspeed < 0.5){
        //     this.currentspeed = 0;
        //     this.movement(event, currentspeed);
        //     return;
        // }
        // this.speed*=this.acceleration;
        // this.movement(event, currentspeed);
        //     return;
        

    }

}




