<script lang="ts" setup>
import {ref} from "vue";
import {Message} from "../Types";
import { PostOllma} from "../../wailsjs/go/main/App";
import {EventsOn} from "../../wailsjs/runtime";
import Logo from "../components/Logo.vue";
import {marked} from 'marked';
import HelloChat from "../Pages/HelloChat.vue";

const TitleSrc = ref("")
const MessageList = ref<Message[]>([])
const InputRef=ref()
const ContentRef=ref()
const ScrollBool=ref(true)
EventsOn("md_get", md =>{
  TitleSrc.value=md.Title
  MessageList.value=md.Data??[]
  console.log(md);
})
const Send=()=>{
  console.log(InputRef.value);
  const NewMessage={
    role:"user",
    content:InputRef.value
  }
  InputRef.value=""
  MessageList.value.push(NewMessage)
  console.log(JSON.stringify(NewMessage));
  PostOllma("llama2",JSON.stringify(MessageList.value)).then(res=>{
    MessageList.value[MessageList.value.length-1]=res
  })
  MessageList.value.push({role:"assistant",content:""})
  const element = ContentRef.value
  element.scrollTop = element.scrollHeight;
}
EventsOn("onMessage",(data)=>{
  console.log(data)
  MessageList.value[MessageList.value.length-1]=data
  if (ScrollBool.value){
    const element = ContentRef.value
    element.scrollTop = element.scrollHeight;
    ScrollBool.value=true
  }
})
const ContentScroll=(e:any)=>{
  const element = e.target;
  ScrollBool.value = element.scrollHeight - element.scrollTop === element.clientHeight;
}
</script>

<template>
  <div class="Box2">
    <div class="TitleBox">
      <div >
        <p>{{ TitleSrc }}</p>
        <img alt="" src="../assets/images/expand_more.svg" v-show="TitleSrc">
      </div>
      <div class="Btn">
        <button>
          <img alt="" src="../assets/images/more_horiz.svg"/>
        </button>
        <button>
          <img alt="" src="../assets/images/share-2.svg"/>
        </button>
      </div>
    </div>
    <div class="ContentBox" v-if="MessageList.length>0" ref="ContentRef" @scroll="(e)=>ContentScroll(e)">
      <div v-for="(item,index) in MessageList" :key="index" class="MessageBox">
        <div v-if="item.role==='user'" class="SendMessage">
          <img
              src="https://s3-alpha-sig.figma.com/img/f6c0/5d02/999762351b16344f2a705a5cff27ddb4?Expires=1711324800&Key-Pair-Id=APKAQ4GOSFWCVNEHN3O4&Signature=Kxv7fhhqNBlR5xR~PLG20ZTOPvZAMLBCBbSUB75dGgQda7DvGdzsYxa-eow3NfVM2xg~KlEU~onwpIHPlaqmpNPWawcFkK5TeY3~ayopha1uJBU1KRIN7HJS07g-RtsdoiMA~X4kSIhub5iA2OJenplcDSBJBW69sMYjeLpCN1dG6DQ6nJcCq3I2GKjNMRerZD31xPjgVrkXi33C92EYiWdJSuUjNkc3pIGSBLdrKhwgEszBYAz627yoXGSNOr2nJz6QjUYwffbRvUA-i2aEOrfuS~9l-vFq2yJa2st6Y6JkaGKhroCRlRrVUcNK3hmVlZAc3UKBhHw6wxd0DcdolA__"/>
          <p v-html="marked.parse(item.content)">
          </p>
        </div>
        <div v-if="item.role==='assistant'" class="RevMessage">
          <div>
            <logo/>
          </div>
          <p v-html="marked.parse(item.content)">
          </p>
        </div>
        <hr class="hr-solid">
      </div>

    </div>
    <HelloChat v-if="MessageList.length<1"/>
    <div class="Folder">
      <div>
        <img alt="" src="../assets/images/paperclip.svg">
        <input placeholder="Type a message" type="text"  v-model="InputRef">
        <button @click="Send" >
          <img alt="" src="../assets/images/send.svg">
        </button>
      </div>
      <button>
        <img alt="" src="../assets/images/sun.svg">
      </button>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.Box2 {
  flex: 1;
  display: flex;
  flex-direction: column;
}

button {
  background: none;
  border: none;
}

.TitleBox {
  height: 64px;
  display: flex;
  padding: 12px;
  justify-content: space-between;
  align-items: center;
  background: rgba(13, 26, 26, 0.80);
  backdrop-filter: blur(24px);

  & > div {
    display: flex;
    align-items: center;
    gap: 12px;

    p {
      color: #F7FAF9;
      font-size: 16px;
      font-style: normal;
      font-weight: 700;
      line-height: 150%; /* 24px */
    }
  }

  .Btn {
    display: flex;
    align-items: center;
    gap: 12px;

    & > button:hover {
      background: #2D3838;

      img {
        filter: drop-shadow(100px 0 0 #14CC8F);
      }
    }

    & > button {
      overflow: hidden;
      height: 48px;
      border-radius: 8px;

      img {
        width: 24px;
        height: 24px;
        margin: 12px;

        filter: drop-shadow(100px 0 0 #7B9999);
      }
    }
  }
}

.ContentBox {
  flex: 1;
  padding: 0 15%;
  margin: 0 5px;
  font-size: 16px;
  line-height: 20px;
  font-weight: bold;
  white-space: pre-wrap;
  overflow-y: auto;

  .SendMessage {
    display: flex;
    align-items: start;
    gap: 12px;

    img {
      height: 30px;
      border-radius: 20px;
    }

    & > div:nth-child(2) {
      margin-top: 4px;
      width: 100%;
    }
  }

  .hr-solid {
    border: 0;
    border-top: 1px double rgba(128, 128, 128, 0.61);
    margin-bottom: 20px;
    margin-top: 20px;
  }

  .RevMessage {
    display: flex;
    align-items: start;
    gap: 12px;

    & > div {
      width: 30px;
    }

    & > p {
      margin-top: 4px;
    }
  }
}

.Folder {
  display: flex;
  height: 88px;
  padding: 16px 8px;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 8px;
  position: relative;

  & > div {
    display: flex;
    width: 60%;
    min-width: 500px;
    padding: 12px 8px;
    align-items: center;
    gap: 8px;
    border-radius: 12px;
    border: 1px solid #2D3838;
    background: #0D1A1A;
    box-shadow: 0 1px 1px 0 rgba(0, 0, 0, 0.05), 0 2px 4px 0 rgba(0, 0, 0, 0.05);

    input {
      flex: 1;
      height: 100%;
      background: transparent;
      outline: transparent;
      border: none;
      color: white;
      font-size: 20px;
    }

    button {
      height: 40px;
      width: 40px;
      display: flex;
      background: #14CC8F;
      border-radius: 8px;
      align-items: center;
      justify-content: center;
    }
  }

  & > button {
    position: absolute;
    width: 40px;
    height: 40px;
    border-radius: 8px;
    display: flex;
    align-items: center;
    justify-content: center;
    border: 1px solid #2D3838;
    background: #0D1A1A;
    box-shadow: 0 1px 1px 0 rgba(0, 0, 0, 0.05), 0 2px 4px 0 rgba(0, 0, 0, 0.05);
    right: 16px;
    bottom: 16px;
  }
}

</style>