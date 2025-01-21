<script>
    import { writable } from 'svelte/store';
    import Login from './Login.svelte';
    import Register from './Register.svelte';

    // Bikin Stores buat manage ganti Page
    const currentPage = writable('login');

    // Ganti Page
    function changePage(event) {
        currentPage.set(event.detail);
    }
</script>

<main>
    {#if $currentPage === 'login'}
        <Login on:changePage={changePage} />
    {:else if $currentPage === 'register'}
        <Register on:changePage={changePage} />
    {/if}
</main>
