# Roadmap: Recipe Manager

## Overview

**Phases:** 4
**Total Requirements:** 28
**Coverage:** 100% of v1 requirements mapped

## Phase 1: Core Recipe CRUD & API Foundation

**Goal:** Build the Go HTTP server with complete recipe CRUD operations, in-memory storage, categories, and Vercel deployment configuration.

**Requirements:** RECV-01, RECV-02, RECV-03, RECV-04, RECV-05, RECV-06, RECV-07

**Success Criteria:**
1. API endpoints for create, read (list + detail), update, and delete recipes return correct JSON responses
2. Recipes store all fields: title, description, ingredients, steps, prep_time, cook_time, servings, tags, image_url, category
3. Recipes can be filtered by category (breakfast, lunch, dinner, dessert, snack)
4. Server starts and handles requests via Vercel-compatible entry point at api/index.go
5. In-memory store persists data within a single server lifecycle

## Phase 2: Search, Scaling & Enhanced Recipe Features

**Goal:** Add ingredient search, keyword search, and recipe scaling capabilities to the API.

**Requirements:** SRCH-01, SRCH-02, SRCH-03, SCAL-01, SCAL-02

**Success Criteria:**
1. User can search recipes by ingredient name and get matching recipe cards
2. User can search recipes by title/keyword and get matching results
3. User can request a scaled version of a recipe with adjusted ingredient quantities
4. Scaling is proportional — doubling servings doubles all ingredient amounts
5. Search and scaling endpoints return proper JSON with all recipe fields

## Phase 3: Meal Planning & Shopping Lists

**Goal:** Build meal plan management (create, assign recipes to days/meals, view calendar) and shopping list generation with ingredient aggregation.

**Requirements:** MEAL-01, MEAL-02, MEAL-03, MEAL-04, MEAL-05, SHOP-01, SHOP-02, SHOP-03, SHOP-04

**Success Criteria:**
1. User can create a weekly meal plan and assign recipes to day+meal slots
2. User can view a meal plan with 7 days x meal types grid structure
3. User can remove recipes from meal plan slots
4. Shopping list aggregates ingredients across all recipes in a meal plan
5. Duplicate ingredients are combined with summed quantities
6. Shopping list items can be checked off (state tracked in-memory)

## Phase 4: Frontend UI

**Goal:** Build the complete frontend with recipe grid, detail view, meal plan calendar, shopping list, search, and add/edit forms — clean, food-focused design.

**Requirements:** FRNT-01, FRNT-02, FRNT-03, FRNT-04, FRNT-05, FRNT-06, FRNT-07

**Success Criteria:**
1. Home page shows recipe card grid with images, titles, and categories
2. Clicking a recipe card opens detail page with full ingredients and steps
3. Meal plan page displays a 7-day calendar grid with assigned recipes
4. Shopping list page shows items with checkboxes that toggle checked state
5. Search page allows filtering by ingredient with results displayed as cards
6. Add/edit recipe form captures all recipe fields with validation
7. UI is clean, responsive, and food-focused with warm color palette

---
*Roadmap created: 2026-03-10*
*Last updated: 2026-03-10 after initial creation*
