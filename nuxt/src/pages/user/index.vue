<script setup>
import { ref } from "@vue/reactivity";
const { data: fetchUsers } = await useFetch(useNuxtApp().$baseUrl("/api/user"));

const firstName = ref("");
const lastName = ref("");
const users = ref(fetchUsers);

const createUser = async () => {
  const postData = {
    firstName: firstName.value,
    lastName: lastName.value,
  };

  try {
    const user = await $fetch("/api/user", {
        method: "POST",
        body: postData,
    });
    firstName.value = "";
    lastName.value = "";
    
    users.value.push(user);
  } catch (e) {
    console.error(e);
  }

};

</script>

<template>
  <div>
    <form @submit.prevent="createUser">
      <input v-model="firstName" placeholder="firstName" />
      <input v-model="lastName" placeholder="lastName" />
      <button type="submit">送信</button>
    </form>

    <table>
      <thead>
        <tr>
          <td>firstName</td>
          <td>lastName</td>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(user, index) in users" :key="index">
          <td>{{ user.firstName }}</td>
          <td>{{ user.lastName }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>