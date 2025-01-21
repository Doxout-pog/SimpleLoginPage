<script>
    import { writable } from 'svelte/store';
    import Login from './Login.svelte';
    import Register from './Register.svelte';

    // Create a writable store to manage the current page
    const currentPage = writable('login');

    // Function to change the current page
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
