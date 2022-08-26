<template>
    <div>
        <form @submit.prevent="submit">
            <input placeholder="firstName" />
            <input placeholder="lastName" />
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

<script lang="ts">
import Vue from 'vue'
import axios from 'axios';
export default Vue.extend({
    data() {
        return {
            firstName: "",
            lastName: "",
            users: [
                {
                    firstName: "huga",
                    lastName: "hoge",
                }
            ],
        }
    },
    async created() {
        const res = await axios.get("/api/hello")
        const data = res.data
        this.users.push(data)
    },
    methods: {
        async submit() {
            const data = {
                firstName: this.firstName,
                lastName: this.lastName,
            };

            const res = await axios.post("/api/users", data);
            console.log(res);
        },
    }
})
</script>