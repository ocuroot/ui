// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.920
package navbar

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

// NavItem interface allows mixing of links and dropdowns in any order
type NavItem interface {
	IsNavItem()
}

// NavLink represents a simple navigation link
type NavLink struct {
	Name   string
	URL    string
	Active bool
}

func (NavLink) IsNavItem() {}

// DropdownItem represents an item within a dropdown menu
type DropdownItem struct {
	Name string
	URL  string
}

// NavDropdown represents a dropdown menu with multiple items
type NavDropdown struct {
	ID    string
	Name  string
	Items []DropdownItem
}

func (NavDropdown) IsNavItem() {}

type NavbarConfig struct {
	BrandName       string
	BrandURL        string
	LogoURL         string
	Items           []NavItem
	ShowThemeToggle bool
	ShowUserMenu    bool
	UserAvatarURL   string
}

func Navbar(config NavbarConfig) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<nav class=\"ocuroot-navbar\"><div class=\"navbar-container\"><div class=\"navbar-brand\"><a href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 templ.SafeURL
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinURLErrs(templ.URL(config.BrandURL))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/navbar/navbar.templ`, Line: 46, Col: 40}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\" class=\"brand-link\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if config.LogoURL != "" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<img src=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(config.LogoURL)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/navbar/navbar.templ`, Line: 48, Col: 31}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\" alt=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(config.BrandName)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/navbar/navbar.templ`, Line: 48, Col: 56}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\" class=\"brand-logo\"> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<span class=\"brand-text\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(config.BrandName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/navbar/navbar.templ`, Line: 50, Col: 48}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "</span></a></div><button class=\"navbar-toggle\" id=\"navbar-toggle\" aria-label=\"Toggle navigation\"><span class=\"hamburger-line\"></span> <span class=\"hamburger-line\"></span> <span class=\"hamburger-line\"></span></button><div class=\"navbar-menu\" id=\"navbar-menu\"><ul class=\"navbar-nav\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, item := range config.Items {
			switch v := item.(type) {
			case NavLink:
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "<li class=\"nav-item\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if v.Active {
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "<a href=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var6 templ.SafeURL
					templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinURLErrs(templ.URL(v.URL))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/navbar/navbar.templ`, Line: 67, Col: 35}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "\" class=\"nav-link active\" aria-current=\"page\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var7 string
					templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(v.Name)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/navbar/navbar.templ`, Line: 67, Col: 90}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "</a>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "<a href=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var8 templ.SafeURL
					templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinURLErrs(templ.URL(v.URL))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/navbar/navbar.templ`, Line: 69, Col: 35}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "\" class=\"nav-link\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var9 string
					templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(v.Name)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/navbar/navbar.templ`, Line: 69, Col: 63}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "</a>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "</li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			case NavDropdown:
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "<li class=\"nav-item dropdown\"><a href=\"#\" class=\"nav-link dropdown-toggle\" data-dropdown=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var10 string
				templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(v.ID)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/navbar/navbar.templ`, Line: 74, Col: 73}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var11 string
				templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(v.Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/navbar/navbar.templ`, Line: 75, Col: 17}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, " <span class=\"dropdown-arrow\">▼</span></a><ul class=\"dropdown-menu\" id=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var12 string
				templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(v.ID)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/navbar/navbar.templ`, Line: 78, Col: 43}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				for _, dropdownItem := range v.Items {
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "<li><a href=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var13 templ.SafeURL
					templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinURLErrs(templ.URL(dropdownItem.URL))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/navbar/navbar.templ`, Line: 80, Col: 51}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "\" class=\"dropdown-link\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var14 string
					templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(dropdownItem.Name)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/navbar/navbar.templ`, Line: 80, Col: 95}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "</a></li>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "</ul></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "</ul><div class=\"navbar-actions\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if config.ShowThemeToggle {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 25, "<button class=\"theme-toggle\" id=\"theme-toggle\" aria-label=\"Toggle theme\"><svg class=\"theme-icon\" width=\"20\" height=\"20\" viewBox=\"0 0 24 24\" fill=\"currentColor\"><!-- Moon icon (dark mode) --><path class=\"moon-icon\" d=\"M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z\"></path><!-- Sun icon (light mode) --><g class=\"sun-icon\" style=\"display: none;\"><circle cx=\"12\" cy=\"12\" r=\"5\"></circle> <path d=\"M12 1v2M12 21v2M4.22 4.22l1.42 1.42M18.36 18.36l1.42 1.42M1 12h2M21 12h2M4.22 19.78l1.42-1.42M18.36 5.64l1.42-1.42\"></path></g></svg></button> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if config.ShowUserMenu {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 26, "<div class=\"user-menu\"><button class=\"user-button\" aria-label=\"User menu\"><img src=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var15 string
			templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(config.UserAvatarURL)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/navbar/navbar.templ`, Line: 105, Col: 39}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 27, "\" alt=\"User\" class=\"user-avatar\"></button></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 28, "</div></div></div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
