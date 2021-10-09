<template>
  <div class="home">
    <!-- <img alt="Vue logo" src="../assets/logo.png"> -->
    <BoardsPage msg="Welcome to Your Tudoo board"/>
  </div>
</template>

<script>
// @ is an alias to /src
import BoardsPage from '@/components/BoardsPage.vue'

export default {
  name: 'Home',
  components: {
    BoardsPage
  }
}
</script>
