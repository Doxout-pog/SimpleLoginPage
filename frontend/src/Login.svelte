<script>
    import { createEventDispatcher } from 'svelte';

    let username = "";
    let password = "";
    let message = "";

    const dispatch = createEventDispatcher();

    // Login Handle
    async function handleLogin() {
        try {
            const response = await fetch("http://localhost:8081/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded", // Encode
                },
                body: new URLSearchParams({
                    username,
                    password,
                }),
            });

            if (response.ok) {
                const result = await response.text();
                message = result; // success message
            } else if (response.status === 401) {
                message = "Invalid username or password.";
            } else {
                message = "An error occurred. Please try again.";
            }
        } catch (error) {
            console.error("Error:", error);
            message = "Unable to connect to the server. Please try again later.";
        }
    }

    // Ganti Page ke Register
    function goToRegister() {
        dispatch('changePage', 'register');
    }
</script>

<style>
    /* Display Form */
    form {
        display: flex;
        flex-direction: column;
        max-width: 300px;
        margin: 20px auto;
    }

    label {
        margin-bottom: 5px;
    }

    input {
        margin-bottom: 15px;
        padding: 8px;
        font-size: 1em;
    }

    button {
        padding: 8px;
        font-size: 1em;
        background-color: #007bff;
        color: white;
        border: none;
        border-radius: 4px;
        cursor: pointer;
    }

    button:hover {
        background-color: #0056b3;
    }

    p {
        text-align: center;
    }

    p.error {
        color: red;
    }

    a {
        color: #007bff;
        cursor: pointer;
    }

    a:hover {
        text-decoration: underline;
    }
</style>

<main>
    <h1>Login</h1>
    <form on:submit|preventDefault={handleLogin}>
        <label for="username">Username</label>
        <input id="username" type="text" bind:value={username} required />

        <label for="password">Password</label>
        <input id="password" type="password" bind:value={password} required />

        <button type="submit">Login</button>
    </form>

    <p class={message === "Login successful" ? "success" : "error"}>{message}</p>

    <p>
        Don't have an account? <a on:click={goToRegister}>Register here</a>
    </p>
</main>
