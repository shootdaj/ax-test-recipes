package frontend

import "net/http"

// ServeIndex serves the main HTML page for the frontend SPA.
func ServeIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(indexHTML))
}

const indexHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Recipe Manager</title>
<style>
*{margin:0;padding:0;box-sizing:border-box}
:root{
  --bg:#fdf6f0;--card-bg:#fff;--primary:#e07a5f;--primary-dark:#c9684f;
  --secondary:#81b29a;--accent:#f2cc8f;--text:#3d405b;--text-light:#6b6e8a;
  --border:#e8ddd4;--shadow:0 2px 8px rgba(61,64,91,0.08);
  --radius:12px;--font:'Segoe UI',system-ui,-apple-system,sans-serif;
}
body{font-family:var(--font);background:var(--bg);color:var(--text);min-height:100vh}
a{color:var(--primary);text-decoration:none}
button{cursor:pointer;font-family:var(--font)}
.btn{
  display:inline-flex;align-items:center;gap:6px;padding:10px 20px;
  border:none;border-radius:8px;font-size:14px;font-weight:600;
  transition:all .2s;
}
.btn-primary{background:var(--primary);color:#fff}
.btn-primary:hover{background:var(--primary-dark);transform:translateY(-1px)}
.btn-secondary{background:var(--secondary);color:#fff}
.btn-secondary:hover{opacity:.9}
.btn-outline{background:transparent;border:2px solid var(--primary);color:var(--primary)}
.btn-outline:hover{background:var(--primary);color:#fff}
.btn-sm{padding:6px 14px;font-size:13px}
.btn-danger{background:#e74c3c;color:#fff}
.btn-danger:hover{background:#c0392b}

/* Navigation */
nav{
  background:#fff;border-bottom:1px solid var(--border);padding:0 24px;
  display:flex;align-items:center;height:64px;position:sticky;top:0;z-index:100;
  box-shadow:var(--shadow);
}
nav .logo{font-size:22px;font-weight:700;color:var(--primary);margin-right:32px}
nav .nav-links{display:flex;gap:4px}
nav .nav-links button{
  padding:8px 16px;border:none;background:transparent;font-size:14px;
  font-weight:500;color:var(--text-light);border-radius:8px;transition:all .2s;
}
nav .nav-links button:hover,nav .nav-links button.active{
  background:var(--primary);color:#fff;
}

/* Layout */
.container{max-width:1200px;margin:0 auto;padding:24px}
.page{display:none}.page.active{display:block}

/* Cards */
.card-grid{display:grid;grid-template-columns:repeat(auto-fill,minmax(280px,1fr));gap:20px}
.recipe-card{
  background:var(--card-bg);border-radius:var(--radius);overflow:hidden;
  box-shadow:var(--shadow);transition:all .2s;cursor:pointer;
}
.recipe-card:hover{transform:translateY(-4px);box-shadow:0 8px 24px rgba(61,64,91,0.12)}
.recipe-card img{width:100%;height:200px;object-fit:cover;background:#e8ddd4}
.recipe-card .card-body{padding:16px}
.recipe-card .card-title{font-size:18px;font-weight:600;margin-bottom:4px}
.recipe-card .card-meta{font-size:13px;color:var(--text-light);display:flex;gap:12px;margin-top:8px}
.recipe-card .card-category{
  display:inline-block;padding:3px 10px;border-radius:20px;font-size:12px;
  font-weight:600;background:var(--accent);color:var(--text);margin-top:8px;
}
.no-img{display:flex;align-items:center;justify-content:center;height:200px;background:var(--accent);font-size:48px}

/* Recipe Detail */
.recipe-detail{background:var(--card-bg);border-radius:var(--radius);box-shadow:var(--shadow);overflow:hidden}
.recipe-detail img{width:100%;max-height:400px;object-fit:cover}
.recipe-detail .detail-body{padding:32px}
.recipe-detail h1{font-size:28px;margin-bottom:8px}
.recipe-detail .meta{display:flex;gap:20px;color:var(--text-light);margin-bottom:20px;flex-wrap:wrap}
.detail-section{margin-top:24px}
.detail-section h2{font-size:18px;margin-bottom:12px;color:var(--primary);border-bottom:2px solid var(--accent);padding-bottom:6px}
.ingredient-list{list-style:none}
.ingredient-list li{padding:8px 0;border-bottom:1px solid var(--border);display:flex;justify-content:space-between}
.steps-list{list-style:none;counter-reset:step}
.steps-list li{
  counter-increment:step;padding:12px 0 12px 40px;position:relative;
  border-bottom:1px solid var(--border);
}
.steps-list li::before{
  content:counter(step);position:absolute;left:0;top:12px;
  width:28px;height:28px;border-radius:50%;background:var(--primary);color:#fff;
  display:flex;align-items:center;justify-content:center;font-size:13px;font-weight:600;
}
.scale-controls{display:flex;align-items:center;gap:12px;margin-top:16px;padding:12px;background:var(--bg);border-radius:8px}
.scale-controls input{width:80px;padding:6px 10px;border:1px solid var(--border);border-radius:6px;font-size:14px}

/* Search */
.search-bar{
  display:flex;gap:12px;margin-bottom:24px;padding:16px;background:var(--card-bg);
  border-radius:var(--radius);box-shadow:var(--shadow);
}
.search-bar input,.search-bar select{
  padding:10px 16px;border:1px solid var(--border);border-radius:8px;font-size:14px;
  font-family:var(--font);flex:1;
}
.search-bar input:focus,.search-bar select:focus{outline:none;border-color:var(--primary)}

/* Meal Plan Calendar */
.calendar{
  display:grid;grid-template-columns:repeat(7,1fr);gap:8px;margin-top:16px;
}
.calendar-day{
  background:var(--card-bg);border-radius:var(--radius);padding:12px;min-height:200px;
  box-shadow:var(--shadow);
}
.calendar-day h3{
  font-size:14px;text-transform:uppercase;color:var(--primary);margin-bottom:12px;
  text-align:center;padding-bottom:8px;border-bottom:2px solid var(--accent);
}
.meal-slot{
  margin-bottom:8px;padding:8px;background:var(--bg);border-radius:8px;font-size:13px;
  position:relative;
}
.meal-slot .meal-label{font-weight:600;color:var(--text-light);font-size:11px;text-transform:uppercase;margin-bottom:2px}
.meal-slot .meal-remove{
  position:absolute;top:4px;right:6px;background:none;border:none;color:#e74c3c;
  font-size:16px;cursor:pointer;opacity:0;transition:opacity .2s;
}
.meal-slot:hover .meal-remove{opacity:1}
.add-meal-btn{
  width:100%;padding:6px;border:1px dashed var(--border);border-radius:6px;
  background:transparent;color:var(--text-light);font-size:12px;cursor:pointer;
  margin-top:4px;
}
.add-meal-btn:hover{border-color:var(--primary);color:var(--primary)}

/* Shopping List */
.shopping-list{background:var(--card-bg);border-radius:var(--radius);padding:24px;box-shadow:var(--shadow)}
.shopping-item{
  display:flex;align-items:center;gap:12px;padding:10px 0;border-bottom:1px solid var(--border);
}
.shopping-item input[type=checkbox]{
  width:20px;height:20px;accent-color:var(--secondary);
}
.shopping-item.checked .item-text{text-decoration:line-through;color:var(--text-light)}
.item-text{font-size:15px;flex:1}
.item-qty{font-size:14px;color:var(--text-light);min-width:80px;text-align:right}

/* Form */
.form-page{background:var(--card-bg);border-radius:var(--radius);padding:32px;box-shadow:var(--shadow)}
.form-group{margin-bottom:20px}
.form-group label{display:block;font-weight:600;margin-bottom:6px;font-size:14px}
.form-group input,.form-group textarea,.form-group select{
  width:100%;padding:10px 14px;border:1px solid var(--border);border-radius:8px;
  font-size:14px;font-family:var(--font);
}
.form-group textarea{min-height:100px;resize:vertical}
.form-group input:focus,.form-group textarea:focus,.form-group select:focus{outline:none;border-color:var(--primary)}
.ingredient-inputs{display:flex;gap:8px}
.ingredient-inputs input{flex:1}
.ingredient-inputs input:first-child{flex:2}
.dynamic-list{margin-top:8px}
.dynamic-item{display:flex;gap:8px;align-items:center;margin-bottom:8px}
.dynamic-item .remove-item{background:none;border:none;color:#e74c3c;font-size:18px;cursor:pointer}

/* Modal */
.modal-overlay{
  position:fixed;inset:0;background:rgba(0,0,0,0.5);display:none;align-items:center;
  justify-content:center;z-index:200;
}
.modal-overlay.show{display:flex}
.modal{background:var(--card-bg);border-radius:var(--radius);padding:24px;max-width:400px;width:90%}
.modal h3{margin-bottom:16px}

/* Responsive */
@media(max-width:768px){
  .calendar{grid-template-columns:1fr}
  .card-grid{grid-template-columns:1fr 1fr}
  .search-bar{flex-direction:column}
}
@media(max-width:480px){
  .card-grid{grid-template-columns:1fr}
  nav .logo{font-size:18px;margin-right:16px}
  nav .nav-links button{padding:6px 10px;font-size:13px}
}
</style>
</head>
<body>

<nav>
  <div class="logo">Recipe Manager</div>
  <div class="nav-links">
    <button class="active" onclick="showPage('recipes')">Recipes</button>
    <button onclick="showPage('search')">Search</button>
    <button onclick="showPage('mealplan')">Meal Plan</button>
    <button onclick="showPage('shopping')">Shopping</button>
  </div>
</nav>

<div class="container">
  <!-- Recipes Page -->
  <div id="page-recipes" class="page active">
    <div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:20px">
      <h1>My Recipes</h1>
      <button class="btn btn-primary" onclick="showPage('add-recipe')">+ Add Recipe</button>
    </div>
    <div style="margin-bottom:16px;display:flex;gap:8px;flex-wrap:wrap">
      <button class="btn btn-sm btn-outline filter-btn active" onclick="filterCategory('','this')">All</button>
      <button class="btn btn-sm btn-outline filter-btn" onclick="filterCategory('breakfast',this)">Breakfast</button>
      <button class="btn btn-sm btn-outline filter-btn" onclick="filterCategory('lunch',this)">Lunch</button>
      <button class="btn btn-sm btn-outline filter-btn" onclick="filterCategory('dinner',this)">Dinner</button>
      <button class="btn btn-sm btn-outline filter-btn" onclick="filterCategory('dessert',this)">Dessert</button>
      <button class="btn btn-sm btn-outline filter-btn" onclick="filterCategory('snack',this)">Snack</button>
    </div>
    <div id="recipe-grid" class="card-grid"></div>
  </div>

  <!-- Recipe Detail Page -->
  <div id="page-detail" class="page">
    <button class="btn btn-outline btn-sm" onclick="showPage('recipes')" style="margin-bottom:16px">&larr; Back</button>
    <div id="recipe-detail" class="recipe-detail"></div>
  </div>

  <!-- Search Page -->
  <div id="page-search" class="page">
    <h1 style="margin-bottom:16px">Search Recipes</h1>
    <div class="search-bar">
      <input type="text" id="search-input" placeholder="Search by title or keyword..." oninput="doSearch()">
      <input type="text" id="ingredient-input" placeholder="Filter by ingredient..." oninput="doSearch()">
    </div>
    <div id="search-results" class="card-grid"></div>
  </div>

  <!-- Add/Edit Recipe Page -->
  <div id="page-add-recipe" class="page">
    <button class="btn btn-outline btn-sm" onclick="showPage('recipes')" style="margin-bottom:16px">&larr; Back</button>
    <div class="form-page">
      <h1 id="form-title">Add Recipe</h1>
      <input type="hidden" id="edit-recipe-id">
      <div class="form-group">
        <label>Title *</label>
        <input type="text" id="recipe-title" placeholder="Recipe name">
      </div>
      <div class="form-group">
        <label>Description</label>
        <textarea id="recipe-desc" placeholder="Brief description..."></textarea>
      </div>
      <div class="form-group">
        <label>Category</label>
        <select id="recipe-category">
          <option value="">Select category</option>
          <option value="breakfast">Breakfast</option>
          <option value="lunch">Lunch</option>
          <option value="dinner">Dinner</option>
          <option value="dessert">Dessert</option>
          <option value="snack">Snack</option>
        </select>
      </div>
      <div class="form-group">
        <label>Image URL</label>
        <input type="url" id="recipe-image" placeholder="https://example.com/photo.jpg">
      </div>
      <div style="display:grid;grid-template-columns:1fr 1fr 1fr;gap:16px">
        <div class="form-group"><label>Prep Time (min)</label><input type="number" id="recipe-prep" min="0" value="0"></div>
        <div class="form-group"><label>Cook Time (min)</label><input type="number" id="recipe-cook" min="0" value="0"></div>
        <div class="form-group"><label>Servings *</label><input type="number" id="recipe-servings" min="1" value="4"></div>
      </div>
      <div class="form-group">
        <label>Tags (comma-separated)</label>
        <input type="text" id="recipe-tags" placeholder="quick, easy, healthy">
      </div>
      <div class="form-group">
        <label>Ingredients *</label>
        <div id="ingredients-list" class="dynamic-list"></div>
        <button class="btn btn-sm btn-outline" onclick="addIngredientRow()" style="margin-top:8px">+ Add Ingredient</button>
      </div>
      <div class="form-group">
        <label>Steps *</label>
        <div id="steps-list" class="dynamic-list"></div>
        <button class="btn btn-sm btn-outline" onclick="addStepRow()" style="margin-top:8px">+ Add Step</button>
      </div>
      <div style="display:flex;gap:12px;margin-top:24px">
        <button class="btn btn-primary" onclick="saveRecipe()">Save Recipe</button>
        <button class="btn btn-outline" onclick="showPage('recipes')">Cancel</button>
      </div>
    </div>
  </div>

  <!-- Meal Plan Page -->
  <div id="page-mealplan" class="page">
    <div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:16px">
      <h1>Meal Plan</h1>
      <div style="display:flex;gap:8px;align-items:center">
        <select id="plan-select" onchange="loadMealPlan(this.value)">
          <option value="">Select a plan...</option>
        </select>
        <button class="btn btn-primary btn-sm" onclick="createMealPlan()">+ New Plan</button>
      </div>
    </div>
    <div id="calendar" class="calendar"></div>
  </div>

  <!-- Shopping List Page -->
  <div id="page-shopping" class="page">
    <div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:16px">
      <h1>Shopping List</h1>
      <div style="display:flex;gap:8px;align-items:center">
        <select id="shopping-plan-select" onchange="loadShoppingList(this.value)">
          <option value="">Select a meal plan...</option>
        </select>
        <button class="btn btn-secondary btn-sm" onclick="loadShoppingList(document.getElementById('shopping-plan-select').value)">Refresh</button>
      </div>
    </div>
    <div id="shopping-list" class="shopping-list"><p style="color:var(--text-light)">Select a meal plan to generate your shopping list.</p></div>
  </div>
</div>

<!-- Assign Recipe Modal -->
<div id="assign-modal" class="modal-overlay" onclick="if(event.target===this)closeModal()">
  <div class="modal">
    <h3>Add Recipe to Meal</h3>
    <input type="hidden" id="modal-day">
    <input type="hidden" id="modal-meal">
    <input type="hidden" id="modal-plan">
    <div class="form-group">
      <label>Recipe</label>
      <select id="modal-recipe-select"></select>
    </div>
    <div style="display:flex;gap:8px;justify-content:flex-end">
      <button class="btn btn-outline btn-sm" onclick="closeModal()">Cancel</button>
      <button class="btn btn-primary btn-sm" onclick="assignRecipe()">Add</button>
    </div>
  </div>
</div>

<script>
const API='/api';
let currentCategory='';
let allRecipes=[];

// --- Navigation ---
function showPage(page){
  document.querySelectorAll('.page').forEach(p=>p.classList.remove('active'));
  document.getElementById('page-'+page).classList.add('active');
  document.querySelectorAll('.nav-links button').forEach(b=>b.classList.remove('active'));
  const navMap={recipes:0,search:1,mealplan:2,shopping:3};
  if(navMap[page]!==undefined)document.querySelectorAll('.nav-links button')[navMap[page]].classList.add('active');
  if(page==='recipes')loadRecipes();
  if(page==='mealplan'){loadMealPlans();const sel=document.getElementById('plan-select');if(sel.value)loadMealPlan(sel.value)}
  if(page==='shopping')loadMealPlansForShopping();
  if(page==='add-recipe'&&!document.getElementById('edit-recipe-id').value)resetForm();
}

// --- Recipes ---
async function loadRecipes(){
  const url=currentCategory?API+'/recipes?category='+currentCategory:API+'/recipes';
  const r=await fetch(url);allRecipes=await r.json();renderRecipeGrid(allRecipes);
}
function renderRecipeGrid(recipes){
  const grid=document.getElementById('recipe-grid');
  if(!recipes||recipes.length===0){grid.innerHTML='<p style="color:var(--text-light);grid-column:1/-1;text-align:center;padding:40px">No recipes yet. Click "+ Add Recipe" to get started!</p>';return}
  grid.innerHTML=recipes.map(r=>{
    const img=r.image_url?'<img src="'+r.image_url+'" alt="'+r.title+'" onerror="this.outerHTML=\'<div class=no-img>&#127859;</div>\'">':'<div class="no-img">&#127859;</div>';
    return '<div class="recipe-card" onclick="viewRecipe(\''+r.id+'\')">'+img+'<div class="card-body"><div class="card-title">'+r.title+'</div>'+(r.description?'<div style="font-size:13px;color:var(--text-light);margin-top:4px">'+r.description.substring(0,80)+'</div>':'')+'<div class="card-meta"><span>&#9201; '+(r.prep_time+r.cook_time)+' min</span><span>&#127869; '+r.servings+' servings</span></div>'+(r.category?'<span class="card-category">'+r.category+'</span>':'')+'</div></div>'
  }).join('');
}
function filterCategory(cat,btn){
  currentCategory=cat;document.querySelectorAll('.filter-btn').forEach(b=>b.classList.remove('active'));
  if(btn&&btn.classList)btn.classList.add('active');loadRecipes();
}

// --- Recipe Detail ---
async function viewRecipe(id){
  const r=await(await fetch(API+'/recipes/'+id)).json();
  const d=document.getElementById('recipe-detail');
  const img=r.image_url?'<img src="'+r.image_url+'" alt="'+r.title+'" onerror="this.style.display=\'none\'">':'';
  d.innerHTML=img+'<div class="detail-body"><div style="display:flex;justify-content:space-between;align-items:flex-start"><div><h1>'+r.title+'</h1>'+(r.category?'<span class="card-category">'+r.category+'</span>':'')+'</div><div style="display:flex;gap:8px"><button class="btn btn-outline btn-sm" onclick="editRecipe(\''+r.id+'\')">Edit</button><button class="btn btn-danger btn-sm" onclick="deleteRecipe(\''+r.id+'\')">Delete</button></div></div>'+(r.description?'<p style="margin-top:12px;color:var(--text-light)">'+r.description+'</p>':'')+'<div class="meta"><span>&#9201; Prep: '+r.prep_time+' min</span><span>&#127859; Cook: '+r.cook_time+' min</span><span>&#127869; '+r.servings+' servings</span></div><div class="scale-controls"><strong>Scale:</strong><input type="number" id="scale-servings" value="'+r.servings+'" min="1"><button class="btn btn-sm btn-secondary" onclick="scaleRecipe(\''+r.id+'\')">Scale</button></div><div class="detail-section"><h2>Ingredients</h2><ul class="ingredient-list" id="ingredient-display">'+r.ingredients.map(i=>'<li><span>'+i.name+'</span><span>'+i.quantity+' '+i.unit+'</span></li>').join('')+'</ul></div><div class="detail-section"><h2>Steps</h2><ol class="steps-list">'+r.steps.map(s=>'<li>'+s+'</li>').join('')+'</ol></div>'+(r.tags&&r.tags.length?'<div class="detail-section"><h2>Tags</h2><div style="display:flex;gap:6px;flex-wrap:wrap">'+r.tags.map(t=>'<span style="padding:4px 12px;background:var(--bg);border-radius:20px;font-size:13px">'+t+'</span>').join('')+'</div></div>':'')+'</div>';
  showPage('detail');
}
async function scaleRecipe(id){
  const servings=document.getElementById('scale-servings').value;
  const r=await(await fetch(API+'/recipes/'+id+'/scale?servings='+servings)).json();
  document.getElementById('ingredient-display').innerHTML=r.ingredients.map(i=>'<li><span>'+i.name+'</span><span>'+i.quantity.toFixed(1)+' '+i.unit+'</span></li>').join('');
}
async function deleteRecipe(id){
  if(!confirm('Delete this recipe?'))return;
  await fetch(API+'/recipes/'+id,{method:'DELETE'});showPage('recipes');
}

// --- Add/Edit Recipe ---
function resetForm(){
  document.getElementById('edit-recipe-id').value='';
  document.getElementById('form-title').textContent='Add Recipe';
  ['recipe-title','recipe-desc','recipe-image','recipe-tags'].forEach(id=>document.getElementById(id).value='');
  document.getElementById('recipe-category').value='';
  document.getElementById('recipe-prep').value='0';
  document.getElementById('recipe-cook').value='0';
  document.getElementById('recipe-servings').value='4';
  document.getElementById('ingredients-list').innerHTML='';
  document.getElementById('steps-list').innerHTML='';
  addIngredientRow();addIngredientRow();addStepRow();addStepRow();
}
function addIngredientRow(name,qty,unit){
  const div=document.createElement('div');div.className='dynamic-item';
  div.innerHTML='<div class="ingredient-inputs"><input type="text" placeholder="Ingredient name" value="'+(name||'')+'"><input type="number" placeholder="Qty" step="0.1" value="'+(qty||'')+'"><input type="text" placeholder="Unit" value="'+(unit||'')+'"></div><button class="remove-item" onclick="this.parentElement.remove()">&#10005;</button>';
  document.getElementById('ingredients-list').appendChild(div);
}
function addStepRow(text){
  const div=document.createElement('div');div.className='dynamic-item';
  div.innerHTML='<input type="text" placeholder="Step description" style="flex:1" value="'+(text||'')+'"><button class="remove-item" onclick="this.parentElement.remove()">&#10005;</button>';
  document.getElementById('steps-list').appendChild(div);
}
async function editRecipe(id){
  const r=await(await fetch(API+'/recipes/'+id)).json();
  document.getElementById('edit-recipe-id').value=id;
  document.getElementById('form-title').textContent='Edit Recipe';
  document.getElementById('recipe-title').value=r.title;
  document.getElementById('recipe-desc').value=r.description||'';
  document.getElementById('recipe-category').value=r.category||'';
  document.getElementById('recipe-image').value=r.image_url||'';
  document.getElementById('recipe-prep').value=r.prep_time||0;
  document.getElementById('recipe-cook').value=r.cook_time||0;
  document.getElementById('recipe-servings').value=r.servings||4;
  document.getElementById('recipe-tags').value=(r.tags||[]).join(', ');
  document.getElementById('ingredients-list').innerHTML='';
  (r.ingredients||[]).forEach(i=>addIngredientRow(i.name,i.quantity,i.unit));
  document.getElementById('steps-list').innerHTML='';
  (r.steps||[]).forEach(s=>addStepRow(s));
  showPage('add-recipe');
}
async function saveRecipe(){
  const id=document.getElementById('edit-recipe-id').value;
  const ingredients=[];
  document.querySelectorAll('#ingredients-list .dynamic-item').forEach(row=>{
    const inputs=row.querySelectorAll('input');
    if(inputs[0].value)ingredients.push({name:inputs[0].value,quantity:parseFloat(inputs[1].value)||0,unit:inputs[2].value});
  });
  const steps=[];
  document.querySelectorAll('#steps-list .dynamic-item input').forEach(input=>{
    if(input.value)steps.push(input.value);
  });
  const tags=document.getElementById('recipe-tags').value.split(',').map(t=>t.trim()).filter(Boolean);
  const recipe={
    title:document.getElementById('recipe-title').value,
    description:document.getElementById('recipe-desc').value,
    category:document.getElementById('recipe-category').value,
    image_url:document.getElementById('recipe-image').value,
    prep_time:parseInt(document.getElementById('recipe-prep').value)||0,
    cook_time:parseInt(document.getElementById('recipe-cook').value)||0,
    servings:parseInt(document.getElementById('recipe-servings').value)||4,
    tags:tags,ingredients:ingredients,steps:steps
  };
  if(!recipe.title||!ingredients.length||!steps.length){alert('Please fill in title, at least one ingredient, and at least one step.');return}
  const method=id?'PUT':'POST';
  const url=id?API+'/recipes/'+id:API+'/recipes';
  const resp=await fetch(url,{method,headers:{'Content-Type':'application/json'},body:JSON.stringify(recipe)});
  if(!resp.ok){const e=await resp.json();alert('Error: '+e.error);return}
  showPage('recipes');
}

// --- Search ---
async function doSearch(){
  const q=document.getElementById('search-input').value;
  const ing=document.getElementById('ingredient-input').value;
  let recipes=[];
  if(ing){const r=await fetch(API+'/search?ingredient='+encodeURIComponent(ing));recipes=await r.json()}
  else if(q){const r=await fetch(API+'/search?q='+encodeURIComponent(q));recipes=await r.json()}
  else{const r=await fetch(API+'/recipes');recipes=await r.json()}
  const grid=document.getElementById('search-results');
  if(!recipes||recipes.length===0){grid.innerHTML='<p style="color:var(--text-light);grid-column:1/-1;text-align:center;padding:40px">No recipes found.</p>';return}
  grid.innerHTML=recipes.map(r=>{
    const img=r.image_url?'<img src="'+r.image_url+'" alt="'+r.title+'" onerror="this.outerHTML=\'<div class=no-img>&#127859;</div>\'">':'<div class="no-img">&#127859;</div>';
    return '<div class="recipe-card" onclick="viewRecipe(\''+r.id+'\')">'+img+'<div class="card-body"><div class="card-title">'+r.title+'</div><div class="card-meta"><span>&#9201; '+(r.prep_time+r.cook_time)+' min</span><span>&#127869; '+r.servings+' servings</span></div>'+(r.category?'<span class="card-category">'+r.category+'</span>':'')+'</div></div>'
  }).join('');
}

// --- Meal Plans ---
async function loadMealPlans(){
  const r=await fetch(API+'/mealplans');const plans=await r.json();
  const sel=document.getElementById('plan-select');const cur=sel.value;
  sel.innerHTML='<option value="">Select a plan...</option>'+plans.map(p=>'<option value="'+p.id+'"'+(p.id===cur?' selected':'')+'>'+p.name+'</option>').join('');
}
async function loadMealPlansForShopping(){
  const r=await fetch(API+'/mealplans');const plans=await r.json();
  const sel=document.getElementById('shopping-plan-select');const cur=sel.value;
  sel.innerHTML='<option value="">Select a meal plan...</option>'+plans.map(p=>'<option value="'+p.id+'"'+(p.id===cur?' selected':'')+'>'+p.name+'</option>').join('');
}
async function createMealPlan(){
  const name=prompt('Meal plan name:');if(!name)return;
  await fetch(API+'/mealplans',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify({name})});
  await loadMealPlans();
  const sel=document.getElementById('plan-select');sel.selectedIndex=sel.options.length-1;
  loadMealPlan(sel.value);
}
async function loadMealPlan(id){
  if(!id){document.getElementById('calendar').innerHTML='';return}
  const plan=await(await fetch(API+'/mealplans/'+id)).json();renderCalendar(plan);
}
function renderCalendar(plan){
  const days=['monday','tuesday','wednesday','thursday','friday','saturday','sunday'];
  const meals=['breakfast','lunch','dinner','snack'];
  const cal=document.getElementById('calendar');
  cal.innerHTML=days.map(day=>{
    const daySlots=plan.slots?plan.slots.filter(s=>s.day.toLowerCase()===day):[];
    return '<div class="calendar-day"><h3>'+day.charAt(0).toUpperCase()+day.slice(1,3)+'</h3>'+meals.map(meal=>{
      const slot=daySlots.find(s=>s.meal_type.toLowerCase()===meal);
      if(slot){
        const recipe=allRecipes.find(r=>r.id===slot.recipe_id);
        return '<div class="meal-slot"><div class="meal-label">'+meal+'</div><div>'+(recipe?recipe.title:'Recipe #'+slot.recipe_id)+'</div><button class="meal-remove" onclick="removeMealSlot(\''+plan.id+'\',\''+day+'\',\''+meal+'\')">&times;</button></div>';
      }
      return '<button class="add-meal-btn" onclick="openAssignModal(\''+plan.id+'\',\''+day+'\',\''+meal+'\')">+ '+meal+'</button>';
    }).join('')+'</div>';
  }).join('');
}
function openAssignModal(planId,day,meal){
  document.getElementById('modal-plan').value=planId;
  document.getElementById('modal-day').value=day;
  document.getElementById('modal-meal').value=meal;
  const sel=document.getElementById('modal-recipe-select');
  sel.innerHTML=allRecipes.map(r=>'<option value="'+r.id+'">'+r.title+'</option>').join('');
  document.getElementById('assign-modal').classList.add('show');
}
function closeModal(){document.getElementById('assign-modal').classList.remove('show')}
async function assignRecipe(){
  const planId=document.getElementById('modal-plan').value;
  const slot={day:document.getElementById('modal-day').value,meal_type:document.getElementById('modal-meal').value,recipe_id:document.getElementById('modal-recipe-select').value};
  await fetch(API+'/mealplans/'+planId+'/slots',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify(slot)});
  closeModal();loadMealPlan(planId);
}
async function removeMealSlot(planId,day,meal){
  await fetch(API+'/mealplans/'+planId+'/slots?day='+day+'&meal_type='+meal,{method:'DELETE'});
  loadMealPlan(planId);
}

// --- Shopping List ---
async function loadShoppingList(planId){
  if(!planId){document.getElementById('shopping-list').innerHTML='<p style="color:var(--text-light)">Select a meal plan to generate your shopping list.</p>';return}
  const items=await(await fetch(API+'/mealplans/'+planId+'/shopping-list')).json();
  const list=document.getElementById('shopping-list');
  if(!items||items.length===0){list.innerHTML='<p style="color:var(--text-light)">No items in shopping list. Add recipes to your meal plan first.</p>';return}
  list.innerHTML=items.map((item,i)=>'<div class="shopping-item" id="shop-item-'+i+'"><input type="checkbox" onchange="toggleShopItem('+i+')"><span class="item-text">'+item.name+'</span><span class="item-qty">'+item.quantity.toFixed(1)+' '+item.unit+'</span></div>').join('');
}
function toggleShopItem(i){
  document.getElementById('shop-item-'+i).classList.toggle('checked');
}

// --- Init ---
loadRecipes();
</script>
</body>
</html>`
