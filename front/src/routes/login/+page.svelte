<script lang="js">
  import axiosApiV1 from "$lib/axios/api/v1/axios";
  import defineToken from "$lib/user/defineUserToken";
  import tokenIsPresent from "$lib/user/tokenIsPresent";

  let user = {
    email: "",
    password: "",
  };

  const handleSubmit = async () => {
    await axiosApiV1
      .post("/users/auth", user)
      .then((res) => {
        defineToken(res.data.token);

        if (tokenIsPresent) {
          window.location.href = "/boards";
        }
      })
      .catch((err) => {
        console.log(err);
      });
  };
</script>

<div class="row m-0 p-0">
  <div
    class="col-lg-6 col-12 bg-gray-light d-flex justify-content-center align-items-center vh100"
  >
    <div class="avatar-login" />
  </div>
  <div
    class="col-12 col-lg-6 d-flex justify-content-center align-items-center flex-column"
  >
    <h1 class="mb-5">Acesse sua conta!</h1>
    <form
      on:submit={handleSubmit}
      class="w100 d-flex justify-content-center align-items-center flex-column"
    >
      <div class="input-login-container w80">
        <span class="text-gray">Email</span>
        <input
          type="text"
          placeholder="insira seu email aqui"
          class="form-control text-gray input-login"
          bind:value={user.email}
        />
      </div>
      <div class="input-login-container w80 mt-4">
        <span class="text-gray">Senha</span>
        <input
          type="password"
          placeholder="insira sua senha aqui"
          class="form-control text-gray input-login"
          bind:value={user.password}
        />
      </div>
      <button class="btn-login w80 mt-5">login</button>

      <div
        class="w100 mt-4 d-flex justify-content-center align-items-center flex-column"
      >
        <div class="w100 d-flex justify-content-center align-items-center">
          <div class="detail-or" />
          <h5 class="color-light-gray">OR</h5>
          <div class="detail-or" />
        </div>
        <h6 class="color-light-gray mt-4">
          Não tem uma conta?
          <a href="/" style="text-decoration: none;">Crie uma!</a>
        </h6>
      </div>
    </form>
  </div>
</div>

<style scoped>
  .vh100 {
    height: 100vh;
  }

  .w100 {
    width: 100%;
  }

  .w80 {
    width: 80%;
  }

  .avatar-login {
    background-image: url(../../assets/imgs/login_avatar.png);
    background-repeat: no-repeat;
    background-size: cover;
    height: 570px;
    width: 570px;
  }

  .bg-gray-light {
    background-color: #e5eff8;
  }

  .text-gray {
    color: #7c838a;
  }

  .input-login {
    padding: 8px 8px;
    width: 100%;
    border-radius: 14px;
    background-color: #e3e5e8;
    font-weight: 200;
    margin-top: 5px;
  }

  .btn-login {
    background-color: #e5eff8;
    border: none;
    border-radius: 14px;
    padding: 10px;
    cursor: pointer;
  }

  .btn-login:hover {
    background-color: #abb4bc;
    transition: 1.5s;
  }

  .color-light-gray {
    color: #b0bac3;
  }

  .detail-or {
    height: 3px;
    width: 3%;
    border-radius: 10px;
    background-color: #b0bac3;
    margin: 0px 20px;
  }
</style>
