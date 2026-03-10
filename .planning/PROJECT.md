# Recipe Manager

## What This Is

A web-based Recipe Manager built with Go (net/http stdlib) that lets users create, organize, and plan meals around their recipes. It provides CRUD operations for recipes, ingredient-based search, weekly meal planning, automatic shopping list generation, recipe scaling, and categorization — all with a clean, food-focused frontend. Deployed as a Vercel serverless function with in-memory storage.

## Core Value

Users can manage their recipes, plan weekly meals, and generate shopping lists from those plans — the complete meal planning workflow in one tool.

## Requirements

### Validated

(None yet — ship to validate)

### Active

- [ ] CRUD for recipes (title, description, ingredients list, steps, prep_time, cook_time, servings, tags, image_url)
- [ ] Ingredient search (find recipes containing specific ingredients)
- [ ] Meal planning (create weekly meal plans assigning recipes to days/meals)
- [ ] Shopping list generation (aggregate ingredients from meal plan)
- [ ] Recipe scaling (adjust ingredient quantities for different serving sizes)
- [ ] Recipe categories (breakfast, lunch, dinner, dessert, snack)
- [ ] Recipe card grid with images
- [ ] Recipe detail page with ingredients/steps
- [ ] Meal plan calendar view (7-day grid)
- [ ] Shopping list with checkboxes
- [ ] Recipe search with ingredient filter
- [ ] Add recipe form
- [ ] Clean, food-focused UI

### Out of Scope

- User authentication — single-user app with in-memory storage
- Persistent database — in-memory storage only for v1
- Image upload — uses image_url field for external image links
- Mobile native app — web-only
- Social/sharing features — personal use tool
- Nutritional information — not in v1

## Context

- Built with Go stdlib net/http (no frameworks)
- Deployed to Vercel as a serverless function via `api/index.go`
- In-memory data storage (no database)
- Frontend served as embedded HTML/CSS/JS from Go handler
- Single binary entry point for Vercel compatibility

## Constraints

- **Tech stack**: Go stdlib only (net/http), no external Go dependencies
- **Deployment**: Vercel serverless — entry point at `api/index.go`
- **Storage**: In-memory only — data resets on cold starts (acceptable for v1)
- **Frontend**: Server-rendered or embedded static HTML/CSS/JS served by Go handler

## Key Decisions

| Decision | Rationale | Outcome |
|----------|-----------|---------|
| Go stdlib net/http | Simplicity, no dependency management, Vercel-compatible | — Pending |
| In-memory storage | Fastest path to working product, no DB setup needed | — Pending |
| Vercel deployment | Free hosting, easy deployment, Go support via @vercel/go | — Pending |
| Embedded frontend | Single binary, no separate build step, simpler deployment | — Pending |

---
*Last updated: 2026-03-10 after initialization*
