# Requirements: Recipe Manager

**Defined:** 2026-03-10
**Core Value:** Users can manage their recipes, plan weekly meals, and generate shopping lists from those plans

## v1 Requirements

### Recipes

- [ ] **RECV-01**: User can create a recipe with title, description, ingredients list, steps, prep_time, cook_time, servings, tags, and image_url
- [ ] **RECV-02**: User can view a list of all recipes as a card grid with images
- [ ] **RECV-03**: User can view a single recipe's full detail (ingredients, steps, times, servings)
- [ ] **RECV-04**: User can update an existing recipe
- [ ] **RECV-05**: User can delete a recipe
- [ ] **RECV-06**: User can assign categories to recipes (breakfast, lunch, dinner, dessert, snack)
- [ ] **RECV-07**: User can filter recipes by category

### Search

- [ ] **SRCH-01**: User can search recipes by ingredient (find recipes containing specific ingredients)
- [ ] **SRCH-02**: User can search recipes by title or keyword
- [ ] **SRCH-03**: Search results display as recipe cards with images

### Scaling

- [ ] **SCAL-01**: User can scale a recipe to a different number of servings
- [ ] **SCAL-02**: Ingredient quantities adjust proportionally when scaling

### Meal Planning

- [ ] **MEAL-01**: User can create a weekly meal plan
- [ ] **MEAL-02**: User can assign recipes to specific days and meals (breakfast, lunch, dinner, snack)
- [ ] **MEAL-03**: User can view meal plan as a 7-day calendar grid
- [ ] **MEAL-04**: User can remove a recipe from a meal plan slot
- [ ] **MEAL-05**: User can view multiple meal plans

### Shopping List

- [ ] **SHOP-01**: User can generate a shopping list from a meal plan
- [ ] **SHOP-02**: Shopping list aggregates identical ingredients across recipes
- [ ] **SHOP-03**: User can check off items on the shopping list
- [ ] **SHOP-04**: Shopping list displays ingredient quantities and units

### Frontend

- [ ] **FRNT-01**: Recipe card grid with images on home page
- [ ] **FRNT-02**: Recipe detail page with ingredients and steps
- [ ] **FRNT-03**: Meal plan calendar view (7-day grid)
- [ ] **FRNT-04**: Shopping list with checkboxes
- [ ] **FRNT-05**: Recipe search with ingredient filter input
- [ ] **FRNT-06**: Add/edit recipe form
- [ ] **FRNT-07**: Clean, food-focused UI design

## v2 Requirements

### Persistence

- **PERS-01**: Data persists across server restarts (database backend)
- **PERS-02**: User accounts and authentication

### Enhanced Features

- **ENHC-01**: Nutritional information per recipe
- **ENHC-02**: Recipe import from URL
- **ENHC-03**: Recipe sharing via link
- **ENHC-04**: Cooking timer integration
- **ENHC-05**: Favorite/bookmark recipes

## Out of Scope

| Feature | Reason |
|---------|--------|
| User authentication | Single-user app for v1, in-memory storage |
| Persistent database | In-memory storage acceptable for v1 demo |
| Image upload | Uses external image URLs instead |
| Mobile native app | Web-first approach |
| Social/sharing features | Personal use tool |
| Nutritional information | Complexity, defer to v2 |

## Traceability

| Requirement | Phase | Status |
|-------------|-------|--------|
| RECV-01 | Phase 1 | Pending |
| RECV-02 | Phase 1 | Pending |
| RECV-03 | Phase 1 | Pending |
| RECV-04 | Phase 1 | Pending |
| RECV-05 | Phase 1 | Pending |
| RECV-06 | Phase 1 | Pending |
| RECV-07 | Phase 1 | Pending |
| SRCH-01 | Phase 2 | Pending |
| SRCH-02 | Phase 2 | Pending |
| SRCH-03 | Phase 2 | Pending |
| SCAL-01 | Phase 2 | Pending |
| SCAL-02 | Phase 2 | Pending |
| MEAL-01 | Phase 3 | Pending |
| MEAL-02 | Phase 3 | Pending |
| MEAL-03 | Phase 3 | Pending |
| MEAL-04 | Phase 3 | Pending |
| MEAL-05 | Phase 3 | Pending |
| SHOP-01 | Phase 3 | Pending |
| SHOP-02 | Phase 3 | Pending |
| SHOP-03 | Phase 3 | Pending |
| SHOP-04 | Phase 3 | Pending |
| FRNT-01 | Phase 4 | Pending |
| FRNT-02 | Phase 4 | Pending |
| FRNT-03 | Phase 4 | Pending |
| FRNT-04 | Phase 4 | Pending |
| FRNT-05 | Phase 4 | Pending |
| FRNT-06 | Phase 4 | Pending |
| FRNT-07 | Phase 4 | Pending |

**Coverage:**
- v1 requirements: 28 total
- Mapped to phases: 28
- Unmapped: 0

---
*Requirements defined: 2026-03-10*
*Last updated: 2026-03-10 after initial definition*
