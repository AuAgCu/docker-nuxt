<template>
    <div>
        <form @submit.prevent="submit">
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

<script lang="ts">
import Vue from 'vue'
import axios from 'axios';

export type User = {
    firstName: string,
    lastName: string,
}

export default Vue.extend({
    data() {
        return {
            firstName: "",
            lastName: "",
            users: Array() ,
        }
    },
    async mounted() {
        const res = await axios.get("/api/user")
        this.users = res.data;
    },
    methods: {
        async submit() {
            const data = {
                firstName: this.firstName,
                lastName: this.lastName,
            };

            try {
                const res = await axios.post("/api/user", data);
                this.users.push(res.data);

                this.firstName = "";
                this.lastName = "";
            } catch (e) {
                console.error(e)
            }
            
        },
    }
})
</script>