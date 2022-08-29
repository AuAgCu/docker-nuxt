<script setup>
    const { data: users } = await useFetch('http://localhost:3000/api/user');
    // console.log(users);
    // const apiFetch = $fetch.create({ baseURL: 'http://localhost:3000' })

    // const {data: users} = await useFetch("/apfi/user");
    console.log(users)
</script>

<template>
    <div>
        <form @submit.prevent="submit">
            <!-- <input v-model="firstName" placeholder="firstName" />
            <input v-model="lastName" placeholder="lastName" /> -->
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