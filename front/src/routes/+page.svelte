<script>
  import "../app.css";

  let email = ''
  let password = ''

</script>

<div class="flex h-screen">
  <div class="flex-1 background-black justify-center items-center flex">
      <div class="flex-1 justify-center items-center flex">
        <div class="w-1/2">
          <form>
            <h1 class="color-white text-5xl font-bold">Login</h1>
            <p class="color-white mt-5 font-extralight">acesse sua conta =)</p>
            <input class="w-full input-login mt-10" type="text" placeholder="email" bind:value={email}/> <br>
            <input class="mt-3 w-full input-login mb-3" type="password" placeholder="password" bind:value={email}/>
            <a href="/" class="color-white mt-5 font-extralight">não tem uma conta? crie uma</a>
            <button class="mt-5 w-full bg-white py-3 rounded-md text-black font-bold" type="submit">Login</button>
          </form>
        </div>

    </div>
  </div>

  <div class="flex-1 background-img background-purple" />
</div>


<style scoped>
.background-img{
  background-image: url("../assets/img/img-login.png");
  background-repeat: no-repeat;
}

.input-login{
  border: none;
  border-bottom: 1px solid white;
  background-color: transparent;
  color:white
}
</style>