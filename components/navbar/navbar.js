/**
 * Ocuroot Navbar - Vanilla JavaScript Implementation
 * Handles mobile menu toggle, dropdown functionality, and theme switching
 */

// Use IIFE to prevent global pollution and redeclaration issues
(function() {
  'use strict';
  
  // Prevent redeclaration if already loaded
  if (typeof window.OcurootNavbar !== 'undefined') {
    return;
  }

class OcurootNavbar {
  constructor() {
    this.init();
  }

  init() {
    this.setupMobileToggle();
    this.setupDropdowns();
    this.setupThemeToggle();
    this.setupClickOutside();
  }

  setupMobileToggle() {
    const toggle = document.getElementById('navbar-toggle');
    const menu = document.getElementById('navbar-menu');

    if (!toggle || !menu) return;

    toggle.addEventListener('click', () => {
      const isActive = toggle.classList.contains('active');
      
      if (isActive) {
        this.closeMobileMenu();
      } else {
        this.openMobileMenu();
      }
    });
  }

  openMobileMenu() {
    const toggle = document.getElementById('navbar-toggle');
    const menu = document.getElementById('navbar-menu');
    
    toggle.classList.add('active');
    menu.classList.add('active');
    
    // Prevent body scroll when mobile menu is open
    document.body.style.overflow = 'hidden';
  }

  closeMobileMenu() {
    const toggle = document.getElementById('navbar-toggle');
    const menu = document.getElementById('navbar-menu');
    
    toggle.classList.remove('active');
    menu.classList.remove('active');
    
    // Restore body scroll
    document.body.style.overflow = '';
    
    // Close any open dropdowns
    this.closeAllDropdowns();
  }

  setupDropdowns() {
    const dropdowns = document.querySelectorAll('.dropdown');
    
    dropdowns.forEach(dropdown => {
      const toggle = dropdown.querySelector('.dropdown-toggle');
      if (!toggle) return;
      
      // Check if we're on a touch device
      const isTouchDevice = 'ontouchstart' in window || navigator.maxTouchPoints > 0;
      
      if (isTouchDevice) {
        // On touch devices, use click to toggle
        toggle.addEventListener('click', (e) => {
          e.preventDefault();
          const isActive = dropdown.classList.contains('active');
          
          // Close all other dropdowns first
          this.closeAllDropdowns();
          
          // Toggle this dropdown
          if (!isActive) {
            dropdown.classList.add('active');
          }
        });
      } else {
        // On non-touch devices, use hover
        dropdown.addEventListener('mouseenter', () => {
          this.closeAllDropdowns();
          dropdown.classList.add('active');
        });
        
        dropdown.addEventListener('mouseleave', () => {
          setTimeout(() => {
            dropdown.classList.remove('active');
          }, 100);
        });
        
        // Still prevent default click behavior
        toggle.addEventListener('click', (e) => {
          e.preventDefault();
        });
      }
    });
  }

  closeAllDropdowns() {
    const activeDropdowns = document.querySelectorAll('.dropdown.active');
    activeDropdowns.forEach(dropdown => {
      dropdown.classList.remove('active');
    });
  }

  setupThemeToggle() {
    const themeToggle = document.getElementById('theme-toggle');
    const themeIcon = themeToggle?.querySelector('.theme-icon');
    
    if (!themeToggle || !themeIcon) return;

    // Get current theme from localStorage or system preference
    const currentTheme = localStorage.getItem('theme') || 
      (window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light');
    
    this.setTheme(currentTheme);

    themeToggle.addEventListener('click', () => {
      const currentTheme = document.documentElement.getAttribute('data-theme') || 'light';
      const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
      this.setTheme(newTheme);
    });
  }

  setTheme(theme) {
    const moonIcon = document.querySelector('.moon-icon');
    const sunIcon = document.querySelector('.sun-icon');
    
    document.documentElement.setAttribute('data-theme', theme);
    localStorage.setItem('theme', theme);
    
    if (moonIcon && sunIcon) {
      if (theme === 'dark') {
        // Show sun icon in dark mode (to switch to light)
        moonIcon.style.display = 'none';
        sunIcon.style.display = 'block';
      } else {
        // Show moon icon in light mode (to switch to dark)
        moonIcon.style.display = 'block';
        sunIcon.style.display = 'none';
      }
    }
  }

  setupClickOutside() {
    document.addEventListener('click', (e) => {
      // Close dropdowns when clicking outside
      if (!e.target.closest('.dropdown')) {
        this.closeAllDropdowns();
      }
      
      // Close mobile menu when clicking outside (on mobile)
      const navbar = e.target.closest('.ocuroot-navbar');
      const isMobile = window.innerWidth <= 768;
      
      if (!navbar && isMobile) {
        this.closeMobileMenu();
      }
    });
    
    // Also handle escape key to close menus
    document.addEventListener('keydown', (e) => {
      if (e.key === 'Escape') {
        this.closeAllDropdowns();
        this.closeMobileMenu();
      }
    });
  }

  // Handle window resize
  handleResize() {
    const isMobile = window.innerWidth <= 768;
    
    if (!isMobile) {
      this.closeMobileMenu();
      document.body.style.overflow = '';
    }
  }
}

// Make class available globally
window.OcurootNavbar = OcurootNavbar;

// Initialize navbar when DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
  const navbar = new OcurootNavbar();
  
  // Handle window resize
  window.addEventListener('resize', () => {
    navbar.handleResize();
  });
});

// Export for potential external use
if (typeof module !== 'undefined' && module.exports) {
  module.exports = OcurootNavbar;
}

})(); // End of IIFE
