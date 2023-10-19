<template>
  <div id="loginRef" class="login">
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
    >
      <swiper-slide v-for="item in theme?.backgroundPhoto" :key="item.url">
        <img :src="'../api/' + item.url" style="width: 100%; height: 100%" />
      </swiper-slide>
    </swiper>
    <div class="login__mask" />
    <div class="login__container">
      <img
        :src="'../api/' + theme?.systemLogo[0]?.url"
        class="login__logo"
        :style="`transform: scale(${transformScale});`"
      />
      <div class="login__box" :style="`transform: scale(${transformScale});`">
        <div class="login__box-left">
          <div class="login__title">
            <img
              :src="'../api/' + theme?.loginViewLogo[0]?.url"
              style="
                width: 120px;
                vertical-align: middle;
                margin-right: 5px;
                margin-bottom: 0.5rem;
              "
            />
            <br />{{ theme?.systemName }}
          </div>
          <el-form
            ref="loginForm"
            class="login__form"
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
        <img class="login__box-right" src="../../assets/login/login_left.svg" />
      </div>
      <BottomInfo
        class="login__bottom-info"
        :style="`transform: translateX(-50%) scale(${transformScale});`"
      />
    </div>
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
import { captcha, getThemeById, findThemeByRoleId } from "@/api/user";
import { checkDB } from "@/api/initdb";
import BottomInfo from "@/view/layout/bottomInfo/bottomInfo.vue";
import { reactive, ref, onMounted, onBeforeUnmount } from "vue";
import { ElMessage } from "element-plus";
import { useRouter, useRoute } from "vue-router";
import { useUserStore } from "@/pinia/modules/user";

const defaultThemes = {
  backgroundPhoto: [
    {
      name: "login_background1.png",
      url: "uploads/file/3c5192e8d553bdf18741ab6f01beb669_20231016153502.png",
      uid: 1697443136505,
      status: "success",
    },
    {
      name: "login_background2.png",
      url: "uploads/file/5bc1da8e95ff528d72882f3db39a7d37_20231016153504.png",
      uid: 1697443136506,
      status: "success",
    },
    {
      name: "login_background3.jpg",
      url: "uploads/file/394c1a5822b3934109d2a078f7500a84_20231016153506.jpg",
      uid: 1697443136507,
      status: "success",
    },
  ],
  systemLogo: [
    {
      name: "logo.png",
      url: "uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20231016153500.png",
      uid: 1697443136501,
      status: "success",
    },
  ],
  loginViewLogo: [
    {
      name: "index-removebg-preview.png",
      url: "uploads/file/83baf2016971cfedc7f12efae74b3576_20231016155906.png",
      uid: 1697443241479,
      status: "success",
    },
  ],
  systemName: "智农飞手",
  ID: 7,
  CreatedAt: "2023-10-16T15:35:10.774+08:00",
  UpdatedAt: "2023-10-16T16:06:01.768+08:00",
  themeName: "默认主题",
  userRoles: "9528",
  isOrNoDefaultTheme: true,
};

const theme = ref({
  backgroundPhoto: [],
  systemLogo: [],
  loginViewLogo: [],
  systemName: "",
});

const transformScale = ref(1);
const setTransformScale = () => {
  transformScale.value = window.screen.height / 1080;
};

const router = useRouter();
const route = useRoute();
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
      if (flag) {
        findThemeByRoleId({
          userRoles: userStore.userInfo.authorityId,
        }).then((res) => {
          console.log(res);
          if (res.code !== 7) {
            Object.assign(theme.value, res.data.retheme);

            theme.value.loginViewLogo =
              theme.value.loginViewLogo === ""
                ? []
                : JSON.parse(theme.value.loginViewLogo);
            theme.value.systemLogo =
              theme.value.systemLogo === ""
                ? []
                : JSON.parse(theme.value.systemLogo);
            theme.value.backgroundPhoto =
              theme.value.backgroundPhoto === ""
                ? []
                : JSON.parse(theme.value.backgroundPhoto);
          } else {
            Object.assign(theme.value, defaultThemes);
          }
          localStorage.setItem("theme", JSON.stringify(theme.value));
        });
      }

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
let resizeObserver = null;
const setTheme = (id = "") => {
  getThemeById({ id: id }).then((res) => {
    if (!res.request) {
      Object.assign(theme.value, res.data);

      theme.value.loginViewLogo =
        theme.value.loginViewLogo === ""
          ? []
          : JSON.parse(theme.value.loginViewLogo);
      theme.value.systemLogo =
        theme.value.systemLogo === "" ? [] : JSON.parse(theme.value.systemLogo);
      theme.value.backgroundPhoto =
        theme.value.backgroundPhoto === ""
          ? []
          : JSON.parse(theme.value.backgroundPhoto);
    } else {
      Object.assign(theme.value, defaultThemes);
    }
    localStorage.setItem("theme", JSON.stringify(theme.value));
  });
};
const getTheme = () => {
  if (route.query.id === undefined) {
    const localStorageTheme = localStorage.getItem("theme");
    if (localStorageTheme) {
      theme.value = JSON.parse(localStorageTheme);
    } else {
      setTheme("");
    }
  } else {
    setTheme(route.query.id);
  }
};
getTheme();
onMounted(() => {
  resizeObserver = new window.ResizeObserver(setTransformScale);
  resizeObserver.observe(document.getElementById("loginRef"));
});
onBeforeUnmount(() => {
  resizeObserver.unobserve(document.getElementById("loginRef"));
});
</script>

<style lang="scss" scoped>
.login {
  height: 100%;
  width: 100%;
  position: relative;
  &__mask {
    left: 0;
    right: 0;
    bottom: 0;
    top: 0;
    background: linear-gradient(rgba(0, 0, 0, 0), rgba(0, 0, 0, 0.3));
    z-index: 2;
    position: absolute;
  }
  &__swiper {
    width: 100%;
    height: 100%;
    position: absolute;
    z-index: 1;
    .swiper-slide {
      width: 100%;
    }
  }
  &__logo {
    position: absolute;
    top: 0.625rem;
    left: 0.625rem;
    width: 120px;
    transform: scale(1.5);
    transform-origin: top left;
  }
  &__container {
    width: 100%;
    height: 100%;
    left: 0;
    position: relative;
    z-index: 3;
  }
  &__box {
    border-radius: 1.875rem;
    padding: 3.125rem;
    box-sizing: border-box;
    position: absolute;
    height: 35rem;
    width: 60rem;
    bottom: 4rem;
    left: 3rem;
    display: flex;
    justify-content: space-between;
    background-color: rgba(255, 255, 255, 0.3);
    transform-origin: bottom left;
  }
  &__box-left {
    width: 45%;
    padding: 0.625rem;
    box-sizing: border-box;
  }
  &__title {
    font-size: 2.8rem;
    color: rgba(255, 255, 255, 0.9);
    text-align: center;
    margin-top: -1rem;
    margin-bottom: 2rem;
  }
  &__box-right {
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
    color: white;
    transform-origin: center;
    :deep(a) {
      color: white;
    }
  }
}
</style>
