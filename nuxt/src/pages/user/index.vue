<script setup>
import { ref } from "@vue/reactivity";

const { data: users } = await useFetch("http://localhost:3000/api/user");
const firstName = ref("");
const lastName = ref("");
const createUser = () => {
  const data = {
    firstName: firstName.value,
    lastName: lastName.value,
  };

  try {
    const {data: user} = useFetch("/api/user", {
        method: "POST",
        body: data,
    });
    console.log(user);
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