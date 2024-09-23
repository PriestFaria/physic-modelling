<template>
  <div id="app">
    <h1>Визуализация движения точки на ободе колеса</h1>
    <div>
      <label for="radius">Радиус:</label>
      <input type="number" v-model="radius" id="radius" />
      <label for="speed">Скорость центра масс:</label>
      <input type="number" v-model="speed" id="speed" />
      <button @click="startAnimation">Начать анимацию</button>
    </div>
    <canvas ref="canvas" width="600" height="400" style="border: 1px solid black;"></canvas>
  </div>
</template>

<script>
export default {
  data() {
    return {
      radius: 1, 
      speed: 1,
      distance: 0,
      animationId: null,
      trajectory: [],
    };
  },
  methods: {
    drawGrid(ctx) {
      ctx.strokeStyle = '#e0e0e0';
      ctx.lineWidth = 1;


      for (let x = 0; x <= 600; x += 20) {
        ctx.beginPath();
        ctx.moveTo(x, 0);
        ctx.lineTo(x, 400);
        ctx.stroke();
      }

      for (let y = 0; y <= 400; y += 20) {
        ctx.beginPath();
        ctx.moveTo(0, y);
        ctx.lineTo(600, y);
        ctx.stroke();
      }

      ctx.strokeStyle = 'black';
      ctx.lineWidth = 2;


      ctx.beginPath();
      ctx.moveTo(0, 200);
      ctx.lineTo(600, 200);
      ctx.stroke();


      ctx.beginPath();
      ctx.moveTo(300, 0);
      ctx.lineTo(300, 400);
      ctx.stroke();
    },
    drawWheel(ctx, x, y) {
      ctx.beginPath();
      ctx.arc(x, y, this.radius*20, 0, 2 * Math.PI);
      ctx.strokeStyle = 'blue';
      ctx.stroke();
    },
    drawPoint(ctx, x, y) {
      ctx.beginPath();
      ctx.arc(x, y, 5, 0, 2 * Math.PI);
      ctx.fillStyle = 'red';
      ctx.fill();
    },
    drawTrajectory(ctx) {
      ctx.beginPath();
      ctx.moveTo(this.trajectory[0]?.x, this.trajectory[0]?.y);
      for (let point of this.trajectory) {
        ctx.lineTo(point.x, point.y);
      }
      ctx.strokeStyle = 'green';
      ctx.stroke();
    },
    draw() {
      const canvas = this.$refs.canvas;
      const ctx = canvas.getContext('2d');
      ctx.clearRect(0, 0, canvas.width, canvas.height);

      this.drawGrid(ctx);

      const wheelCenterX = this.distance + this.radius*20;  
      const wheelCenterY = 200;

      this.drawWheel(ctx, wheelCenterX, wheelCenterY);


      const angle = this.distance / (this.radius*20); 
      const pointX = wheelCenterX + (this.radius*20) * Math.cos(angle);
      const pointY = wheelCenterY + (this.radius * 20) * Math.sin(angle);
      this.drawPoint(ctx, pointX, pointY);


      this.trajectory.push({ x: pointX, y: pointY });
      this.drawTrajectory(ctx);
    },
    startAnimation() {
      if (this.animationId) {
        cancelAnimationFrame(this.animationId);
      }
      this.distance = 0; 
      this.trajectory = []; 
      const animate = () => {
        this.distance += this.speed; 
        this.draw();
        this.animationId = requestAnimationFrame(animate);
      };
      animate();
    },
  },
};
</script>

<style>
#app {
  text-align: center;
}
</style>
