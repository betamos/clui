What is a reasonable task? What belongs to a task?

Task with progress might be a subclass of no progress.

Task without progress:

spinner := spinner.New("downloading files").Show()
defer spinner.Abort()
..
spinner.Update("verifying hashes")

if err != nil {
    spinner.Fail(fmt.Sprintf("err: %v", err))
} else {
    spinner.Success("file download successful")
}

# Multi spinner

s1 := spinner.New("building")
s2 := spinner.New("provisioning VMs")

s3 := spinner.Multi("deploying app bacon", s1, s2).Show()
defer s3.Abort()

go func() {
    s1.Update()
}() // thread 1

wg.Wait()
s3.Success("deployment successful")


# Multi spinner

s1 := spinner.New("building")
s2 := spinner.New("provisioning VMs") // two active ones, automatically shown
// Hijack stdout/stderr and show it once released?
defer s2.Abort() // if you're doing tricky stuff, use this one

go func() {
    s1.Update()
    s1.Success()
}() // thread 1

s2.Fail(I failed) // It know to release display due to global state

wg.Wait()

# Spinner with warnings?

s1 := spinner.New("doin me a trick")
s1.Warn("why oh why")





# For general UI:

clui.Colors = False

clui.Theme = ...  // Sensible default

clui.SuccessFile = os.Stdout
clui.FailFile = os.Stderr
clui.WarnFile = os.Stdout
clui.UpdateFile = os.DevNull  // This is unnecessary

# Done

clui.Success("greeted user")
clui.Fail("the file %v had the wrong permissions", f)
clui.Warn("disk almost full")

# Easy progress
task := clui.NewTask("uploading files")
// count to a million
task.Success() // persists something hey

# Multiple processes
t1 := clui.NewTask("uploading files")
// 1s of work
t2 := clui.NewTask("provisioning VMs")
// 2s of work
t2.Fail("provisioning VMs failed")
// 1s of work
t1.Success("upload complete")  // stop displaying these things


# internal data structure

var []Task tasks
