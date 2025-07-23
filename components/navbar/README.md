# Ocuroot Navbar Component

A clean, responsive navbar component built with vanilla CSS and JavaScript - no external dependencies like Tailwind or Flowbite.

## Features

- **Responsive Design**: Horizontal layout on desktop, hamburger menu on mobile
- **Vanilla Implementation**: Pure CSS and JavaScript, no external frameworks
- **Dropdown Menus**: Support for nested navigation items
- **Theme Toggle**: Built-in dark/light mode switching
- **Accessible**: Proper ARIA labels and keyboard navigation support
- **Smooth Animations**: CSS transitions for all interactions

## Usage

### 1. Include the Component

```go
import "github.com/ocuroot/ui/components/navbar"

// In your templ file:
@navbar.Navbar(navbar.NavbarConfig{
    BrandName: "Ocuroot",
    BrandURL:  "/",
    LogoURL:   "/static/logo.svg",
    Links: []navbar.NavLink{
        {Name: "Environments", URL: "/environments", Active: false},
        {Name: "Repositories", URL: "/repositories", Active: true}, // Active page
        {Name: "Commits", URL: "/commits", Active: false},
    },
    Dropdowns: []navbar.Dropdown{
        {
            ID:   "releases",
            Name: "Releases",
            Items: []navbar.DropdownItem{
                {Name: "Releases", URL: "/releases"},
                {Name: "Phases", URL: "/phases"},
            },
        },
    },
    ShowThemeToggle: true,
    ShowUserMenu:    true,
    UserAvatarURL:   "/static/anon_user.svg",
})
```

### 2. Include CSS and JavaScript

Add to your HTML head:

```html
<link rel="stylesheet" href="/components/navbar/navbar.css">
<script src="/components/navbar/navbar.js"></script>
```

### 3. CSS Variables

The navbar uses CSS custom properties that can be customized:

```css
:root {
  --navbar-bg: #242526;        /* Navbar background color */
  --navbar-text: #f0f0f0;      /* Navbar text color */
  --border-color: #e5e7eb;     /* Border color */
}
```

## Structure

```
navbar/
├── navbar.templ    # Templ component
├── navbar.css      # Vanilla CSS styles
├── navbar.js       # Vanilla JavaScript functionality
└── README.md       # This file
```

## Customization

### Adding New Navigation Items

Edit the `navbar.templ` file and add new `nav-item` elements:

```html
<li class="nav-item">
    <a href="/your-page" class="nav-link">Your Page</a>
</li>
```

### Adding New Dropdown Menus

```html
<li class="nav-item dropdown">
    <a href="#" class="nav-link dropdown-toggle" data-dropdown="your-dropdown">
        Your Menu
        <span class="dropdown-arrow">▼</span>
    </a>
    <ul class="dropdown-menu" id="your-dropdown">
        <li><a href="/item1" class="dropdown-link">Item 1</a></li>
        <li><a href="/item2" class="dropdown-link">Item 2</a></li>
    </ul>
</li>
```

## Browser Support

- Chrome/Edge 60+
- Firefox 55+
- Safari 12+
- Mobile browsers with CSS Grid support

## Migration from Old Navbar

To replace the existing Flowbite/Tailwind navbar:

1. Replace `@Navbar()` calls with `@navbar.Navbar(c)`
2. Include the new CSS and JS files
3. Remove Flowbite dependencies
4. Update any custom styling to use the new CSS classes
