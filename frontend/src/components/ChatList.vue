<script lang="ts" setup>
import {onMounted, ref} from "vue";
import { useRouter } from "vue-router";
import {CreateMd, GetAllFile, GetMd} from "../../wailsjs/go/main/App";
import {EventsOn} from "../../wailsjs/runtime";
const router = useRouter();
const UserList=ref<any[any]>([])
const UserListRef=ref()
const SelectedTime=ref()
const OnlyOneNewChat=ref(false)
onMounted(() =>{
  GetAllFile("chats").then((res)=>{
    UserList.value=res
    console.log(res)
  })
})
EventsOn("CreateMd",()=>{
  const User=UserList.value[0][0]
  CreateMd("chats", User.Title,User.Time,User.Desc)
})
const AddChat = () => {
  if (OnlyOneNewChat.value===false){
    const User={
      Title: "New Chat",
      Time: Date.now(),
      Desc: "# New Chat"
    }
    if(!UserList.value[0]){
      UserList.value[0]=[User]
    }else {
      UserList.value[0]=[User,...UserList.value[0]]
    }
    UserListRef.value.scrollTop = 0;
    OnlyOneNewChat.value=true
    SelectedTime.value=User.Time
    GetMd("chats",User.Time)
  }
}
const SwitchChat=(Time:number)=>{
  SelectedTime.value=Time
  GetMd("chats",Time)
}
</script>


<template>
  <div class="Box4">
    <div class="Header">
      <div>
        <p>Chats</p>
        <button @click="AddChat">
          <img src="../assets/images/plus.svg" alt=""/>
        </button>
      </div>
      <div>
        <img src="../assets/images/search.svg" alt="">
        <input placeholder="Search"/>
        <img src="../assets/images/sort.svg" alt="">
      </div>
    </div>
    <div class="Chats" ref="UserListRef">
      <div v-for="(item,index) in UserList" :key="index">
        <p v-if="index==0&&item">ToDay</p>
        <p v-if="index==1&&item">Yesterday</p>
        <p v-if="index==2&&item">Previous 7 days</p>
        <p v-if="index==3&&item">Previous 30 days</p>
        <div class="Chat" v-for="(item2,index) in item" :key="index" @click="()=>SwitchChat(item2.Time)" :class="{'Selected':item2.Time==SelectedTime}">
          <img src="../assets/images/messages.svg" alt="">
          <div>
            <p>{{item2.Title}}</p>
            <p>{{item2.Desc}}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.Box4 {
  width: 351px;
  background: rgba(31, 38, 38, 0.80);
  height: 100%;
  display: flex;
  flex-direction: column;
}
.Header{
  height: 128px;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 16px;
  &>div:nth-child(1){
    display: flex;
    align-items: center;
    justify-content: space-between;
    p{
      color: var(--Text-Default-Default, #F7FAF9);
      font-size: 20px;
      font-style: normal;
      font-weight: 700;
      line-height: normal;
    }
    button{
      display: flex;
      height: var(--Size-size-6, 40px);
      padding: 0 var(--Spacing-1, 8px);
      justify-content: center;
      align-items: center;
      gap: var(--Spacing-spacing-0, 0px);
      border-radius: 8px;
      border: 1px solid var(--Border-default, #2D3838);
      background: var(--Background-Default-Default, #0D1A1A);
      box-shadow: 0 1px 1px 0 rgba(0, 0, 0, 0.05), 0 2px 4px 0 rgba(0, 0, 0, 0.05);
    }
  }
  &>div:nth-child(2){
    display: flex;
    height: 40px;
    padding: var(--Spacing-1, 8px);
    align-items: center;
    gap: 8px;
    align-self: stretch;
    border-radius: var(--Spacing-1, 8px);
    border: 1px solid var(--Border-default, #2D3838);
    background: var(--Background-Default-Default, #0D1A1A);
    img{
      display: flex;
      width: 24px;
      height: 24px;
      justify-content: center;
      align-items: center;
    }
    input{
      background: none;
      outline: none;
      border: none;
      flex: 1;
      color: #A3BFBA;
      font-size: 14px;
      font-style: normal;
      font-weight: 400;
      line-height: 150%;
    }
  }
}
.Selected{
  background: #13a26f;
}
.Chats{
  padding: 0 12px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  overflow-y: auto;

  &>div{
    display: flex;
    flex-direction: column;
    gap: 12px;
    padding: 0 16px;
    &>p{
      color: #A3BFBA;
      font-size: 12px;
      font-style: normal;
      font-weight: 500;
      line-height: normal;
    }
    .Chat{
      display: flex;
      gap: 12px;
      align-items: center;
      padding: 8px;
      border-radius: 8px;
      img{
        width: 32px;
        height: 32px;
      }
      &>div{
        &>p:nth-child(1){
          overflow: hidden;
          color: #F7FAF9;
          text-overflow: ellipsis;
          font-size: 14px;
          font-style: normal;
          font-weight: 500;
          line-height: 150%; /* 21px */
        }
        &>p:nth-child(2){
          color: #A3BFBA;
          font-size: 12px;
          font-style: normal;
          font-weight: 500;
          line-height: normal;
        }
      }
    }
  }
}
</style>