import os
import subprocess

# Ask the user for Go code
print("Please enter your Go code below. When done, type 'END' on a new line and press Enter.")
code_lines = []
while True:
    line = input()
    if line == "END":
        break
    code_lines.append(line)

# Combine the lines of code into a single string
code = "\n".join(code_lines)

# Save code to a temporary file
filename = "temp_code.go"
with open(filename, 'w') as file:
    file.write(code)

try:
    # Compile the Go code
    compile_process = subprocess.run(
        ["go", "build", "-o", "temp_program", filename],
        capture_output=True, text=True
    )

    # Check if there was a compilation error
    if compile_process.returncode != 0:
        print("Compilation Error:")
        print(compile_process.stderr)
    else:
        # Run the compiled program
        run_process = subprocess.run(
            ["./temp_program"],
            capture_output=True, text=True
        )
        
        # Show output or runtime error
        if run_process.returncode == 0:
            print("Output:")
            print(run_process.stdout)
        else:
            print("Runtime Error:")
            print(run_process.stderr)

except Exception as e:
    print("An error occurred:", e)

finally:
    # Clean up files
    if os.path.exists(filename):
        os.remove(filename)
    if os.path.exists("temp_program"):
        os.remove("temp_program")
