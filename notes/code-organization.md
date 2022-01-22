# Code Organization

- Some codebase are easy to navigate
- Others are hard to follow
- Code organization is often the difference

A good code structrue will:
1. Make it easier to guess where bugs are
2. Make it easier to add new features
3. Add more ...

## Flat Structure

All code is in a single package
Files can still separate code:

```
myapp/ 
    gallery_store.go
    gallery_handler.go
    gallery_templates.go
    user_store.go
    user_hander.go
    router.go
    ...

```

## Separation of concerns

Separate code based on duties

HTML & CSS is traditional example:
- HTML focuses on overall structrue
- CSS focuses on styling it

Model-View-Controller (MVC) is a popular structure following this strategy. In a web app:
- `models` ==> data, logic, rules, usually database
- `views` ==> rendering things, usually HTML
- `controller` ==> connects it all, accept user input, passes that to models to do stuff, then passes data to views to render things; usually handlers

```bash
mapp/
    controllers/
        user_hander.go
        gallery_handler.go
        ...
    views/
        user_templates.go
        ...
    models/

Doesn't need to be named `models`, `views`, and `controllers`.

[Buffalo](https://gobuffalo.io/en) uses a variation of MVC, but doesn't name directories exactly. Instead

Very common in framwork because it is predictable. Make it easier for a framwork to guess where centain files will be

Ruby on Rails, Django, etc use MVC


