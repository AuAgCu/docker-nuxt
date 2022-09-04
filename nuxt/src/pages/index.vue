<script setup>

    const config = useRuntimeConfig();
    const { data: hoge } = $fetch(`${config.BASE_URL}/api/jwt`);
    // console.log(hoge);
    // console.log(hoge.value)
</script>

<template>
  <div>
    <h1>Hello, World</h1>
    <NuxtLink to="/user">ユーザー一覧</NuxtLink>
    <NuxtLink to="/hello">Python</NuxtLink>
  </div>
</template>