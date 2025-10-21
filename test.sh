#!/bin/bash

# Citizen API Test Script
# Tests all API endpoints with sample data

BASE_URL="http://localhost:8089/api/v1"
EMAIL="aidosg65@gmail.com"
PASSWORD="Test123456!"

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Test counter
PASSED=0
FAILED=0

# Helper function to print test results
print_test() {
    echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo -e "${YELLOW}Testing: $1${NC}"
}

print_success() {
    echo -e "${GREEN}✓ PASS:${NC} $1"
    ((PASSED++))
}

print_error() {
    echo -e "${RED}✗ FAIL:${NC} $1"
    ((FAILED++))
}

print_response() {
    echo -e "${BLUE}Response:${NC}"
    echo "$1" | jq '.' 2>/dev/null || echo "$1"
}

# Function to extract value from JSON
extract_json() {
    echo "$1" | jq -r "$2" 2>/dev/null
}

echo -e "${BLUE}"
echo "╔════════════════════════════════════════════════════════════╗"
echo "║          CITIZEN API ENDPOINT TEST SUITE                  ║"
echo "╚════════════════════════════════════════════════════════════╝"
echo -e "${NC}"

# ============================================
# AUTH ENDPOINTS
# ============================================

echo -e "\n${YELLOW}═══════════════ AUTH ENDPOINTS ═══════════════${NC}\n"

# Test 1: Register
print_test "POST /auth/register - Register new user"
REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/register" \
  -H "Content-Type: application/json" \
  -d "{
    \"first_name\": \"Test\",
    \"last_name\": \"User\",
    \"email\": \"$EMAIL\",
    \"password\": \"$PASSWORD\",
    \"confirm_password\": \"$PASSWORD\",
    \"birth_date\": \"1995-05-15T00:00:00Z\"
  }")

TOKEN=$(extract_json "$REGISTER_RESPONSE" '.token')
if [ ! -z "$TOKEN" ] && [ "$TOKEN" != "null" ]; then
    print_success "User registered successfully"
    print_response "$REGISTER_RESPONSE"
else
    print_error "Registration failed"
    print_response "$REGISTER_RESPONSE"
fi

# Test 2: Login
print_test "POST /auth/login - Login user"
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d "{
    \"email\": \"$EMAIL\",
    \"password\": \"$PASSWORD\"
  }")

TOKEN=$(extract_json "$LOGIN_RESPONSE" '.token')
if [ ! -z "$TOKEN" ] && [ "$TOKEN" != "null" ]; then
    print_success "Login successful"
    print_response "$LOGIN_RESPONSE"
else
    print_error "Login failed"
    print_response "$LOGIN_RESPONSE"
    exit 1
fi

# Test 3: Get Profile
print_test "GET /auth/profile - Get user profile"
PROFILE_RESPONSE=$(curl -s -X GET "$BASE_URL/auth/profile" \
  -H "Authorization: Bearer $TOKEN")

USER_ID=$(extract_json "$PROFILE_RESPONSE" '.profile.id')
if [ ! -z "$USER_ID" ] && [ "$USER_ID" != "null" ]; then
    print_success "Profile retrieved successfully"
    print_response "$PROFILE_RESPONSE"
else
    print_error "Failed to get profile"
    print_response "$PROFILE_RESPONSE"
fi

# Test 4: Forget Password
print_test "POST /auth/forget-password - Request password reset"
FORGET_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/forget-password" \
  -H "Content-Type: application/json" \
  -d "{
    \"email\": \"$EMAIL\"
  }")

ATTEMPT_ID=$(extract_json "$FORGET_RESPONSE" '.attempt_id')
if [ ! -z "$ATTEMPT_ID" ] && [ "$ATTEMPT_ID" != "null" ]; then
    print_success "Password reset OTP sent to email"
    print_response "$FORGET_RESPONSE"
    echo -e "${YELLOW}Note: Check email $EMAIL for OTP code${NC}"
else
    print_error "Failed to request password reset"
    print_response "$FORGET_RESPONSE"
fi

# Test 5: Update Password
print_test "PUT /auth/password - Update password"
NEW_PASSWORD="NewTest123456!"
UPDATE_PASSWORD_RESPONSE=$(curl -s -X PUT "$BASE_URL/auth/password" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d "{
    \"old_password\": \"$PASSWORD\",
    \"password\": \"$NEW_PASSWORD\",
    \"confirm_password\": \"$NEW_PASSWORD\"
  }")

if echo "$UPDATE_PASSWORD_RESPONSE" | grep -q "profile"; then
    print_success "Password updated successfully"
    print_response "$UPDATE_PASSWORD_RESPONSE"
    PASSWORD="$NEW_PASSWORD"
else
    print_error "Failed to update password"
    print_response "$UPDATE_PASSWORD_RESPONSE"
fi

# Test 6: Update Email
print_test "PUT /auth/email - Update email"
NEW_EMAIL="new_$EMAIL"
UPDATE_EMAIL_RESPONSE=$(curl -s -X PUT "$BASE_URL/auth/email" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d "{
    \"email\": \"$NEW_EMAIL\",
    \"password\": \"$PASSWORD\"
  }")

if echo "$UPDATE_EMAIL_RESPONSE" | grep -q "$NEW_EMAIL"; then
    print_success "Email updated successfully"
    print_response "$UPDATE_EMAIL_RESPONSE"
    EMAIL="$NEW_EMAIL"
else
    print_error "Failed to update email"
    print_response "$UPDATE_EMAIL_RESPONSE"
fi

# ============================================
# CATEGORY ENDPOINTS
# ============================================

echo -e "\n${YELLOW}═══════════════ CATEGORY ENDPOINTS ═══════════════${NC}\n"

# Test 7: Create Category
print_test "POST /category/ - Create category"
CATEGORY_RESPONSE=$(curl -s -X POST "$BASE_URL/category/" \
  -H "Content-Type: application/json" \
  -d "{
    \"name\": \"Education\",
    \"description\": \"Educational benefits and scholarships\"
  }")

CATEGORY_ID=$(extract_json "$CATEGORY_RESPONSE" '.category.id')
if [ ! -z "$CATEGORY_ID" ] && [ "$CATEGORY_ID" != "null" ]; then
    print_success "Category created successfully"
    print_response "$CATEGORY_RESPONSE"
else
    print_error "Failed to create category"
    print_response "$CATEGORY_RESPONSE"
fi

# Test 8: Create Second Category
print_test "POST /category/ - Create second category"
CATEGORY2_RESPONSE=$(curl -s -X POST "$BASE_URL/category/" \
  -H "Content-Type: application/json" \
  -d "{
    \"name\": \"Healthcare\",
    \"description\": \"Healthcare benefits and medical services\"
  }")

CATEGORY2_ID=$(extract_json "$CATEGORY2_RESPONSE" '.category.id')
if [ ! -z "$CATEGORY2_ID" ] && [ "$CATEGORY2_ID" != "null" ]; then
    print_success "Second category created successfully"
    print_response "$CATEGORY2_RESPONSE"
else
    print_error "Failed to create second category"
    print_response "$CATEGORY2_RESPONSE"
fi

# Test 9: List Categories
print_test "POST /category/list - List all categories"
CATEGORIES_LIST=$(curl -s -X POST "$BASE_URL/category/list" \
  -H "Content-Type: application/json" \
  -d "{
    \"limit\": 10,
    \"offset\": 0,
    \"search\": \"\"
  }")

TOTAL_CATEGORIES=$(extract_json "$CATEGORIES_LIST" '.total')
if [ ! -z "$TOTAL_CATEGORIES" ] && [ "$TOTAL_CATEGORIES" != "null" ]; then
    print_success "Categories listed (total: $TOTAL_CATEGORIES)"
    print_response "$CATEGORIES_LIST"
else
    print_error "Failed to list categories"
    print_response "$CATEGORIES_LIST"
fi

# Test 10: Get Category
print_test "GET /category/$CATEGORY_ID - Get category by ID"
GET_CATEGORY=$(curl -s -X GET "$BASE_URL/category/$CATEGORY_ID")

if echo "$GET_CATEGORY" | grep -q "Education"; then
    print_success "Category retrieved successfully"
    print_response "$GET_CATEGORY"
else
    print_error "Failed to get category"
    print_response "$GET_CATEGORY"
fi

# Test 11: Update Category
print_test "PUT /category/$CATEGORY_ID - Update category"
UPDATE_CATEGORY=$(curl -s -X PUT "$BASE_URL/category/$CATEGORY_ID" \
  -H "Content-Type: application/json" \
  -d "{
    \"name\": \"Updated Education\",
    \"description\": \"Updated educational benefits\"
  }")

if echo "$UPDATE_CATEGORY" | grep -q "Updated Education"; then
    print_success "Category updated successfully"
    print_response "$UPDATE_CATEGORY"
else
    print_error "Failed to update category"
    print_response "$UPDATE_CATEGORY"
fi

# ============================================
# FILTER ENDPOINTS
# ============================================

echo -e "\n${YELLOW}═══════════════ FILTER ENDPOINTS ═══════════════${NC}\n"

# Test 12: Create Filter (Age Range)
print_test "POST /filter/ - Create age range filter"
FILTER_RESPONSE=$(curl -s -X POST "$BASE_URL/filter/" \
  -H "Content-Type: application/json" \
  -d "{
    \"name\": \"Age Range\",
    \"hint\": \"Enter your age range\",
    \"type\": \"NUMBER_RANGE\",
    \"values\": []
  }")

FILTER_ID=$(extract_json "$FILTER_RESPONSE" '.filter.id')
if [ ! -z "$FILTER_ID" ] && [ "$FILTER_ID" != "null" ]; then
    print_success "Filter created successfully"
    print_response "$FILTER_RESPONSE"
else
    print_error "Failed to create filter"
    print_response "$FILTER_RESPONSE"
fi

# Test 13: Create Filter (Status)
print_test "POST /filter/ - Create status filter"
FILTER2_RESPONSE=$(curl -s -X POST "$BASE_URL/filter/" \
  -H "Content-Type: application/json" \
  -d "{
    \"name\": \"Student Status\",
    \"hint\": \"Select your student status\",
    \"type\": \"STRING_RANGE\",
    \"values\": [\"student\", \"graduate\", \"undergrad\"]
  }")

FILTER2_ID=$(extract_json "$FILTER2_RESPONSE" '.filter.id')
if [ ! -z "$FILTER2_ID" ] && [ "$FILTER2_ID" != "null" ]; then
    print_success "Second filter created successfully"
    print_response "$FILTER2_RESPONSE"
else
    print_error "Failed to create second filter"
    print_response "$FILTER2_RESPONSE"
fi

# Test 14: List Filters
print_test "GET /filter/ - List all filters"
FILTERS_LIST=$(curl -s -X GET "$BASE_URL/filter/")

if echo "$FILTERS_LIST" | grep -q "filters"; then
    print_success "Filters listed successfully"
    print_response "$FILTERS_LIST"
else
    print_error "Failed to list filters"
    print_response "$FILTERS_LIST"
fi

# ============================================
# BENEFIT ENDPOINTS
# ============================================

echo -e "\n${YELLOW}═══════════════ BENEFIT ENDPOINTS ═══════════════${NC}\n"

# Test 15: Create Benefit
print_test "POST /benefit/ - Create benefit with filters and categories"
BENEFIT_RESPONSE=$(curl -s -X POST "$BASE_URL/benefit/" \
  -H "Content-Type: application/json" \
  -d "{
    \"title\": \"Student Discount Program\",
    \"content\": \"Get 20% discount on all educational materials\",
    \"bonus\": \"Extra 10% during exam season\",
    \"video_url\": \"https://example.com/video.mp4\",
    \"source_url\": \"https://example.com/benefits/student\",
    \"filters\": [
      {
        \"filter_id\": $FILTER_ID,
        \"from\": \"18\",
        \"to\": \"25\"
      },
      {
        \"filter_id\": $FILTER2_ID,
        \"value\": \"student\"
      }
    ],
    \"categories\": [$CATEGORY_ID, $CATEGORY2_ID]
  }")

BENEFIT_ID=$(extract_json "$BENEFIT_RESPONSE" '.benefit.id')
if [ ! -z "$BENEFIT_ID" ] && [ "$BENEFIT_ID" != "null" ]; then
    print_success "Benefit created successfully"
    print_response "$BENEFIT_RESPONSE"
else
    print_error "Failed to create benefit"
    print_response "$BENEFIT_RESPONSE"
fi

# Test 16: Create Benefit without filters
print_test "POST /benefit/ - Create benefit without filters"
BENEFIT2_RESPONSE=$(curl -s -X POST "$BASE_URL/benefit/" \
  -H "Content-Type: application/json" \
  -d "{
    \"title\": \"General Citizen Benefit\",
    \"content\": \"Available for all citizens\",
    \"bonus\": \"No special requirements\",
    \"filters\": [],
    \"categories\": [$CATEGORY_ID]
  }")

BENEFIT2_ID=$(extract_json "$BENEFIT2_RESPONSE" '.benefit.id')
if [ ! -z "$BENEFIT2_ID" ] && [ "$BENEFIT2_ID" != "null" ]; then
    print_success "Benefit without filters created successfully"
    print_response "$BENEFIT2_RESPONSE"
else
    print_error "Failed to create benefit without filters"
    print_response "$BENEFIT2_RESPONSE"
fi

# Test 17: Get Benefit
print_test "GET /benefit/$BENEFIT_ID - Get benefit by ID"
GET_BENEFIT=$(curl -s -X GET "$BASE_URL/benefit/$BENEFIT_ID")

if echo "$GET_BENEFIT" | grep -q "Student Discount"; then
    print_success "Benefit retrieved successfully"
    print_response "$GET_BENEFIT"
else
    print_error "Failed to get benefit"
    print_response "$GET_BENEFIT"
fi

# Test 18: List Benefits (no filters)
print_test "POST /benefit/list - List all benefits"
LIST_BENEFITS=$(curl -s -X POST "$BASE_URL/benefit/list" \
  -H "Content-Type: application/json" \
  -d "{
    \"limit\": 10,
    \"offset\": 0,
    \"search\": \"\",
    \"filters\": []
  }")

TOTAL_BENEFITS=$(extract_json "$LIST_BENEFITS" '.total')
if [ ! -z "$TOTAL_BENEFITS" ] && [ "$TOTAL_BENEFITS" != "null" ]; then
    print_success "Benefits listed (total: $TOTAL_BENEFITS)"
    print_response "$LIST_BENEFITS"
else
    print_error "Failed to list benefits"
    print_response "$LIST_BENEFITS"
fi

# Test 19: List Benefits with matching filter
print_test "POST /benefit/list - List benefits with matching filter"
LIST_FILTERED=$(curl -s -X POST "$BASE_URL/benefit/list" \
  -H "Content-Type: application/json" \
  -d "{
    \"limit\": 10,
    \"offset\": 0,
    \"search\": \"\",
    \"filters\": [
      {
        \"filter_id\": $FILTER2_ID,
        \"value\": \"student\"
      }
    ]
  }")

FILTERED_TOTAL=$(extract_json "$LIST_FILTERED" '.total')
if [ ! -z "$FILTERED_TOTAL" ] && [ "$FILTERED_TOTAL" != "null" ]; then
    print_success "Filtered benefits listed (total: $FILTERED_TOTAL)"
    print_response "$LIST_FILTERED"
else
    print_error "Failed to list filtered benefits"
    print_response "$LIST_FILTERED"
fi

# Test 20: List Benefits with search
print_test "POST /benefit/list - Search benefits"
SEARCH_BENEFITS=$(curl -s -X POST "$BASE_URL/benefit/list" \
  -H "Content-Type: application/json" \
  -d "{
    \"limit\": 10,
    \"offset\": 0,
    \"search\": \"Student\",
    \"filters\": []
  }")

SEARCH_TOTAL=$(extract_json "$SEARCH_BENEFITS" '.total')
if [ ! -z "$SEARCH_TOTAL" ] && [ "$SEARCH_TOTAL" != "null" ]; then
    print_success "Search results found (total: $SEARCH_TOTAL)"
    print_response "$SEARCH_BENEFITS"
else
    print_error "Failed to search benefits"
    print_response "$SEARCH_BENEFITS"
fi

# Test 21: Update Benefit
print_test "PUT /benefit/$BENEFIT_ID - Update benefit"
UPDATE_BENEFIT=$(curl -s -X PUT "$BASE_URL/benefit/$BENEFIT_ID" \
  -H "Content-Type: application/json" \
  -d "{
    \"title\": \"Updated Student Discount\",
    \"content\": \"Get 25% discount on all educational materials\",
    \"bonus\": \"Extra 15% during exam season\",
    \"filters\": [
      {
        \"filter_id\": $FILTER2_ID,
        \"value\": \"student\"
      }
    ],
    \"categories\": [$CATEGORY_ID]
  }")

if echo "$UPDATE_BENEFIT" | grep -q "Updated Student"; then
    print_success "Benefit updated successfully"
    print_response "$UPDATE_BENEFIT"
else
    print_error "Failed to update benefit"
    print_response "$UPDATE_BENEFIT"
fi

# ============================================
# CLEANUP TESTS
# ============================================

echo -e "\n${YELLOW}═══════════════ CLEANUP TESTS ═══════════════${NC}\n"

# Test 22: Delete Benefit
print_test "DELETE /benefit/$BENEFIT_ID - Delete benefit"
DELETE_BENEFIT=$(curl -s -X DELETE "$BASE_URL/benefit/$BENEFIT_ID")

if echo "$DELETE_BENEFIT" | grep -q "true"; then
    print_success "Benefit deleted successfully"
    print_response "$DELETE_BENEFIT"
else
    print_error "Failed to delete benefit"
    print_response "$DELETE_BENEFIT"
fi

# Test 23: Delete Second Benefit
print_test "DELETE /benefit/$BENEFIT2_ID - Delete second benefit"
DELETE_BENEFIT2=$(curl -s -X DELETE "$BASE_URL/benefit/$BENEFIT2_ID")

if echo "$DELETE_BENEFIT2" | grep -q "true"; then
    print_success "Second benefit deleted successfully"
    print_response "$DELETE_BENEFIT2"
else
    print_error "Failed to delete second benefit"
    print_response "$DELETE_BENEFIT2"
fi

# Test 24: Delete Category
print_test "DELETE /category/$CATEGORY_ID - Delete category"
DELETE_CATEGORY=$(curl -s -X DELETE "$BASE_URL/category/$CATEGORY_ID")

if echo "$DELETE_CATEGORY" | grep -q "true"; then
    print_success "Category deleted successfully"
    print_response "$DELETE_CATEGORY"
else
    print_error "Failed to delete category"
    print_response "$DELETE_CATEGORY"
fi

# Test 25: Delete Second Category
print_test "DELETE /category/$CATEGORY2_ID - Delete second category"
DELETE_CATEGORY2=$(curl -s -X DELETE "$BASE_URL/category/$CATEGORY2_ID")

if echo "$DELETE_CATEGORY2" | grep -q "true"; then
    print_success "Second category deleted successfully"
    print_response "$DELETE_CATEGORY2"
else
    print_error "Failed to delete second category"
    print_response "$DELETE_CATEGORY2"
fi

# Test 26: Delete Profile (last test)
print_test "DELETE /auth/profile - Delete user account"
DELETE_PROFILE=$(curl -s -X DELETE "$BASE_URL/auth/profile" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d "{
    \"password\": \"$PASSWORD\"
  }")

if echo "$DELETE_PROFILE" | grep -q "true"; then
    print_success "User account deleted successfully"
    print_response "$DELETE_PROFILE"
    echo -e "${YELLOW}Note: Deletion confirmation email sent to $EMAIL${NC}"
else
    print_error "Failed to delete user account"
    print_response "$DELETE_PROFILE"
fi

# ============================================
# SUMMARY
# ============================================

echo -e "\n${BLUE}"
echo "╔════════════════════════════════════════════════════════════╗"
echo "║                    TEST SUMMARY                            ║"
echo "╚════════════════════════════════════════════════════════════╝"
echo -e "${NC}"

TOTAL=$((PASSED + FAILED))
echo -e "${GREEN}Passed: $PASSED${NC}"
echo -e "${RED}Failed: $FAILED${NC}"
echo -e "Total:  $TOTAL"
echo ""

if [ $FAILED -eq 0 ]; then
    echo -e "${GREEN}🎉 All tests passed!${NC}"
    exit 0
else
    echo -e "${RED}❌ Some tests failed. Please check the output above.${NC}"
    exit 1
fi
