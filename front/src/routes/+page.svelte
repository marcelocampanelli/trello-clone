<script>
  import axiosInstance from '../middlewares/axios'


  let email = ''
  let password = ''

  const handleSubmit = async () => {
    const data = {
      email: email,
      password: password,
    }

    await axiosInstance
      .post('/users/auth', data)
      .then(response => {
          localStorage.setItem('token', response.data.token)
        },
      ).catch(error => {
        console.log(error)
      })
  }

  $: fieldsNotNull = email === '' || password === ''
</script>

<div class='flex h-screen'>
  <div class='flex-1 background-black justify-center items-center flex'>
    <div class='flex-1 justify-center items-center flex'>
      <div class='w-1/2'>
        <form on:submit={handleSubmit}>
          <h1 class='color-white text-5xl font-bold'>Login</h1>
          <p class='color-white mt-5 font-extralight'>acesse sua conta =)</p>
          <input
            bind:value={email}
            class='w-full input-login mt-10'
            placeholder='email'
            type='text'
          /> <br>
          <input
            bind:value={password}
            class='mt-3 w-full input-login mb-3'
            placeholder='password'
            type='password' />
          <a
            class='color-white mt-5 font-extralight'
            href='/'>
            não tem uma conta? crie uma
          </a>
          <button
            class='mt-5 w-full background-purple  py-3 rounded-md text-black font-bold color-white disabled:opacity-50'
            disabled={fieldsNotNull}
            type='submit'>
            Login
          </button>
        </form>
      </div>
    </div>
  </div>
  <div class='flex-1 background-img background-purple flex flex-1 justify-center pt-20'>
    <div class='flex-1 justify-center items-center flex'>
      <div class='w-1/2'>
        <h1 class='color-white text-5xl font-bold'>Bem vindo</h1>
        <p class='color-white mb-20 font-extralight'>Aqui você pode criar e gerenciar suas tarefas</p>
      </div>
    </div>
  </div>
</div>

<style scoped>
  .background-img {
    background-image: url("../assets/img/img-login.png");
    background-repeat: no-repeat;
  }

  .input-login {
    border: none;
    border-bottom: 1px solid white;
    background-color: transparent;
    color: white;
    padding: 4px;
  }

  .input-login:focus {
    outline-style: none;
  }
</style>
