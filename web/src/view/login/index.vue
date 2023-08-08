<template>
  <div class="login">
    <swiper
      class="login__swiper"
      :modules="[EffectFade, Autoplay]"
      :autoplay="{
        delay: 3000,
      }"
      :loop="true"
      :speed="1500"
      effect="fade"
      :slides-per-view="1"
      @swiper="onSwiper"
      @slideChange="onSlideChange"
    >
      <swiper-slide>
        <img
          src="../../assets/login/login_background1.png"
          style="width: 100%; height: 100%"
        />
      </swiper-slide>
      <swiper-slide>
        <img
          src="../../assets/login/login_background2.png"
          style="width: 100%; height: 100%"
      /></swiper-slide>
      <swiper-slide>
        <img
          src="../../assets/login/login_background3.jpg"
          style="width: 100%; height: 100%"
      /></swiper-slide>
    </swiper>
    <img src="../../assets/login/logo.png" class="login__logo" />
    <div class="login__layout"></div>
    <div class="login__container">
      <div class="login__container-left">
        <div class="login__title">
          <img
            src="../../assets/login/znslogo.svg"
            style="width: 100px; vertical-align: middle; margin-right: 5px"
          />智农飞手
        </div>
        <el-form
          class="login__form"
          ref="loginForm"
          :model="loginFormData"
          :rules="rules"
          :validate-on-rule-change="false"
          @keyup.enter="submitForm"
        >
          <el-form-item prop="username">
            <el-input
              v-model="loginFormData.username"
              size="large"
              placeholder="请输入用户名"
              suffix-icon="user"
            />
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model="loginFormData.password"
              show-password
              size="large"
              type="password"
              placeholder="请输入密码"
            />
          </el-form-item>
          <el-form-item v-if="loginFormData.openCaptcha" prop="captcha">
            <div class="vPicBox">
              <el-input
                v-model="loginFormData.captcha"
                placeholder="请输入验证码"
                size="large"
                style="flex: 1; padding-right: 20px"
              />
              <div class="vPic">
                <img
                  v-if="picPath"
                  :src="picPath"
                  alt="请输入验证码"
                  @click="loginVerify()"
                />
              </div>
            </div>
          </el-form-item>
          <el-form-item>
            <!-- <el-button
              type="primary"
              style="width: 46%"
              size="large"
              @click="checkInit"
            >前往初始化</el-button> -->
            <el-button
              type="primary"
              size="large"
              style="width: 100%"
              @click="submitForm"
              >登 录</el-button
            >
          </el-form-item>
        </el-form>
      </div>
      <img
        class="login__container-right"
        src="../../assets/login/login_left.svg"
      />
    </div>
    <BottomInfo class="login__bottom-info" />
  </div>
</template>

<script>
export default {
  name: "Login",
};
</script>

<script setup>
import { EffectFade, Autoplay } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import "swiper/css";
import "swiper/css/effect-fade";
import "swiper/css/autoplay";
import { captcha } from "@/api/user";
import { checkDB } from "@/api/initdb";
import BottomInfo from "@/view/layout/bottomInfo/bottomInfo.vue";
import { reactive, ref } from "vue";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router";
import { useUserStore } from "@/pinia/modules/user";

const onSwiper = (swiper) => {
  console.log(swiper);
};
const onSlideChange = () => {
  console.log("slide change");
};

const router = useRouter();
// 验证函数
const checkUsername = (rule, value, callback) => {
  if (value.length < 5) {
    return callback(new Error("请输入正确的用户名"));
  } else {
    callback();
  }
};
const checkPassword = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error("请输入正确的密码"));
  } else {
    callback();
  }
};

// 获取验证码
const loginVerify = () => {
  captcha({}).then(async (ele) => {
    rules.captcha.push({
      max: ele.data.captchaLength,
      min: ele.data.captchaLength,
      message: `请输入${ele.data.captchaLength}位验证码`,
      trigger: "blur",
    });
    picPath.value = ele.data.picPath;
    loginFormData.captchaId = ele.data.captchaId;
    loginFormData.openCaptcha = ele.data.openCaptcha;
  });
};
loginVerify();

// 登录相关操作
const loginForm = ref(null);
const picPath = ref("");
const loginFormData = reactive({
  username: "",
  password: "",
  captcha: "",
  captchaId: "",
  openCaptcha: false,
});
const rules = reactive({
  username: [{ validator: checkUsername, trigger: "blur" }],
  password: [{ validator: checkPassword, trigger: "blur" }],
  captcha: [
    {
      message: "验证码格式不正确",
      trigger: "blur",
    },
  ],
});

const userStore = useUserStore();
const login = async () => {
  return await userStore.LoginIn(loginFormData);
};
const submitForm = () => {
  loginForm.value.validate(async (v) => {
    if (v) {
      const flag = await login();
      if (!flag) {
        loginVerify();
      }
    } else {
      ElMessage({
        type: "error",
        message: "请正确填写登录信息",
        showClose: true,
      });
      loginVerify();
      return false;
    }
  });
};

// 跳转初始化
const checkInit = async () => {
  const res = await checkDB();
  if (res.code === 0) {
    if (res.data?.needInit) {
      userStore.NeedInit();
      router.push({ name: "Init" });
    } else {
      ElMessage({
        type: "info",
        message: "已配置数据库信息，无法初始化",
      });
    }
  }
};
</script>

<style lang="scss" scoped>
.login {
  height: 100%;
  width: 100%;
  position: relative;
  &__layout {
    left: 0;
    right: 0;
    bottom: 0;
    top: 0;
    background: linear-gradient(rgba(0, 0, 0, 0), rgba(0, 0, 0, 0.3));
    z-index: 1;
    position: absolute;
  }
  &__swiper {
    width: 100%;
    height: 100%;
    .swiper-slide {
      width: 100%;
    }
  }
  &__logo {
    position: absolute;
    z-index: 2;
    top: 0.625rem;
    left: 0.625rem;
    width: 120px;
  }
  &__container {
    border-radius: 1.875rem;
    padding: 3.125rem;
    box-sizing: border-box;
    position: absolute;
    height: 35rem;
    width: 60rem;
    bottom: 4rem;
    left: 3rem;
    z-index: 2;
    display: flex;
    justify-content: space-between;
    background-color: rgba(255, 255, 255, 0.3);
  }
  &__container-left {
    width: 45%;
    padding: 0.625rem;
    box-sizing: border-box;
  }
  &__title {
    font-size: 3.125rem;
    color: rgba(255, 255, 255, 0.9);
    text-align: center;
    margin-top: 2.5rem;
    margin-bottom: 4rem;
  }
  &__container-right {
  }
  &__form {
    :deep(.el-form-item) {
      margin-bottom: 2.25rem;
    }
  }
  &__bottom-info {
    position: absolute;
    bottom: 0;
    left: 50%;
    transform: translateX(-50%);
    color: white;
    z-index: 2;
    :deep(a) {
      color: white;
    }
  }
}
</style>
