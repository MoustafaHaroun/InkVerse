<script>
  import "../app.css";
</script>

<header id="main-header" class="bg-gray-800 text-white fixed w-full top-0 z-50 translate-all duration-300">
    <div class="container mx-auto px-4">
        <div class="flex items-center justify-between py-4">
            <div class="flex items-center">
                <a href="/"><h1 class="text-xl font-bold">InkVerse</h1></a>
            </div>
            <button id="menu-toggle" class="md:hidden focus:outline-none z-50">
                <svg id="menu-icon" class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7"></path>
                </svg>
                <svg id="close-icon" class="w-6 h-6 hidden" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
            </button>
            <nav id="main-nav" class="fixed inset-y-0 right-0 transform translate-x-full md:relative md:translate-x-0 bg-gray-800 md:bg-transparent w-64 md:w-auto h-full md:h-auto overflow-y-auto md:overflow-visible transition-transform duration-300 ease-in-out md:transition-none">
                <ul class="pt-16 md:pt-0 px-4 md:px-0 md:flex space-y-2 md:space-y-0 md:space-x-4">
                    <li><a href="/login" class="block py-2 md:py-0 hover:text-gray-300 transition duration-200">Login</a></li>
                </ul>
            </nav>
        </div>
    </div>
</header>

<slot />