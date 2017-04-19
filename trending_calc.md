Trending Projects Calculation
=============================

- Get all projects that have had a "note" created in the last month 
- AND where the "note" was created by a user (aka notes.system is false)
- AND where project is visible (aka projects.visibility_level = #{Gitlab::VisibilityLevel::PUBLIC})
- ORDER BY projects with the most notes created in the last month
- LIMIT to the first 100 projects