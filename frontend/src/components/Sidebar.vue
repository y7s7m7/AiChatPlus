<script lang="ts" setup>

import Logo from "./Logo.vue";
import {ref} from "vue";
import {useRouter} from "vue-router";
import UserIcon from "./UserIcon.vue";

const Buttons = [{
  title: "聊天",
  path: "/chat",
},{
  title: "发现",
  path: "/compass",
}, {
  title: "文件",
  path: "/folder",
}, {
  title: "收藏",
  path: "/bookmark",
}, {
  title: "用户",
  path: "/users",
}, {
  title: "更多",
  path: "/more",
}
]
const router=useRouter();
const selectBtn =ref<string>("聊天")
const selectBtns = (title:string) => {
  if (selectBtn.value !== title){
    selectBtn.value=title
  }
  router.push({path:Buttons.find(item => item.title === title)?.path??"/chat"})
}

</script>

<template>
  <div class="SideBox">
    <div class="Btn">
      <Logo/>
      <button @click="selectBtns('聊天')" :class="{'selected':selectBtn==='聊天'}">
        <img src="../assets/images/message-2.svg" alt="">
      </button><button @click="selectBtns('发现')" :class="{'selected':selectBtn==='发现'}">
        <img src="../assets/images/compass.svg" alt="">
      </button><button @click="selectBtns('文件')" :class="{'selected':selectBtn==='文件'}">
        <img src="../assets/images/folder.svg" alt="">
      </button><button @click="selectBtns('收藏')" :class="{'selected':selectBtn==='收藏'}">
        <img src="../assets/images/bookmark.svg" alt="">
      </button><button @click="selectBtns('用户')" :class="{'selected':selectBtn==='用户'}">
        <img src="../assets/images/users.svg" alt="">
      </button><button @click="selectBtns('更多')" :class="{'selected':selectBtn==='更多'}">
        <img src="../assets/images/more_horiz.svg" alt="">
      </button>
    </div>
    <div class="folder">
      <button>
        <img alt="" src="./../assets/images/bell.svg">
      </button>
      <img
          alt=""
          src="https://s3-alpha-sig.figma.com/img/f6c0/5d02/999762351b16344f2a705a5cff27ddb4?Expires=1711324800&Key-Pair-Id=APKAQ4GOSFWCVNEHN3O4&Signature=Kxv7fhhqNBlR5xR~PLG20ZTOPvZAMLBCBbSUB75dGgQda7DvGdzsYxa-eow3NfVM2xg~KlEU~onwpIHPlaqmpNPWawcFkK5TeY3~ayopha1uJBU1KRIN7HJS07g-RtsdoiMA~X4kSIhub5iA2OJenplcDSBJBW69sMYjeLpCN1dG6DQ6nJcCq3I2GKjNMRerZD31xPjgVrkXi33C92EYiWdJSuUjNkc3pIGSBLdrKhwgEszBYAz627yoXGSNOr2nJz6QjUYwffbRvUA-i2aEOrfuS~9l-vFq2yJa2st6Y6JkaGKhroCRlRrVUcNK3hmVlZAc3UKBhHw6wxd0DcdolA__"/>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.SideBox {
  color: white;
  height: 100%;
  width: 80px;
  padding: 16px;
  display: flex;
  flex-direction: column;
  background: #1F2626;
}

button {
  background: none;
  border: none;
}

.Btn {
  display: flex;
  flex: 1;
  justify-content: start;
  align-items: center;
  flex-direction: column;
  gap: 16px;
  .selected{
    background: #2D3838;
    img{
      filter: drop-shadow(100px 0 0 #14CC8F);
    }
  }
  button {
    overflow: hidden;
    height: 48px;
    border-radius: 8px;
    img {
      width: 24px;
      height: 24px;
      margin: 12px;
      transform: translate(-100px);
      filter: drop-shadow(100px 0 0 #7B9999);
    }
  }
}
.folder {
  display: flex;
  flex-direction: column;
  gap: 16px;

  img {
    border-radius: 100px;
  }
}
</style>