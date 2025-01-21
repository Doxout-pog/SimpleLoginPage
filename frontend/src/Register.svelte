<script>
    import { createEventDispatcher } from 'svelte';

    let username = "";
    let password = "";
    let retypePassword = "";
    let message = "";

    const dispatch = createEventDispatcher();

    async function handleRegister() {
        if (password.length < 8) {
            message = "Password must be at least 8 characters long.";
            return;
        }

        if (password !== retypePassword) {
            message = "Passwords do not match.";
            return;
        }

        try {
            const response = await fetch("http://localhost:8081/register", {
                method: "POST",
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                },
                body: new URLSearchParams({
                    username,
                    password,
                }),
            });

            if (response.ok) {
                const result = await response.text();
                message = "Registration successful! Redirecting to login...";
                setTimeout(() => dispatch('changePage', 'login'), 2000);
            } else if (response.status === 409) {
                message = "Username already exists.";
            } else {
                message = "Registration failed. Please try again.";
            }
        } catch (error) {
            console.error("Error:", error);
            message = "Unable to connect to the server. Please try again later.";
        }
    }

    function goToLogin() {
        dispatch('changePage', 'login');
    }
</script>

<style>
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

    p.success {
        color: green;
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
    <h1>Register</h1>
    <form on:submit|preventDefault={handleRegister}>
        <label for="username">Username</label>
        <input id="username" type="text" bind:value={username} required />

        <label for="password">Password</label>
        <input 
            id="password" 
            type="password" 
            bind:value={password} 
            required 
            minlength="8" 
            title="Password must be at least 8 characters long."
        />

        <label for="retypePassword">Retype Password</label>
        <input 
            id="retypePassword" 
            type="password" 
            bind:value={retypePassword} 
            required 
            title="Please re-enter the password to confirm."
        />

        <button type="submit">Register</button>
    </form>

    <p class={message.startsWith("Registration successful") ? "success" : "error"}>{message}</p>

    <p>
        Already have an account? <a on:click={goToLogin}>Login here</a>
    </p>
</main>
