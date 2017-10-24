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


