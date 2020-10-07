package present

import (
	"strconv"

	"github.com/concourse/concourse/atc"
)

func DashboardJob(
	teamName string,
	job atc.DashboardJob,
	pathBuilder PathBuilder,
) atc.Job {
	var presentedNextBuild, presentedFinishedBuild, presentedTransitionBuild *atc.Build

	if job.NextBuild != nil {
		presented := DashboardBuild(*job.NextBuild, pathBuilder)
		presentedNextBuild = &presented
	}

	if job.FinishedBuild != nil {
		presented := DashboardBuild(*job.FinishedBuild, pathBuilder)
		presentedFinishedBuild = &presented
	}

	if job.TransitionBuild != nil {
		presented := DashboardBuild(*job.TransitionBuild, pathBuilder)
		presentedTransitionBuild = &presented
	}

	sanitizedInputs := []atc.JobInput{}
	for _, input := range job.Inputs {
		sanitizedInputs = append(sanitizedInputs, atc.JobInput{
			Name:     input.Name,
			Resource: input.Resource,
			Passed:   input.Passed,
			Trigger:  input.Trigger,
		})
	}

	return atc.Job{
		ID: job.ID,

		Name:                 job.Name,
		PipelineID:           job.PipelineID,
		PipelineName:         job.PipelineName,
		PipelineInstanceVars: job.PipelineInstanceVars,
		TeamName:             teamName,
		Paused:               job.Paused,
		HasNewInputs:         job.HasNewInputs,

		Inputs:  sanitizedInputs,
		Outputs: job.Outputs,

		Groups: job.Groups,

		FinishedBuild:   presentedFinishedBuild,
		NextBuild:       presentedNextBuild,
		TransitionBuild: presentedTransitionBuild,
	}
}

func DashboardBuild(build atc.DashboardBuild, pathBuilder PathBuilder) atc.Build {
	apiURL, err := pathBuilder.CreatePathForRoute(atc.GetBuild, map[string]string{
		"build_id":  strconv.Itoa(build.ID),
		"team_name": build.TeamName,
	})
	if err != nil {
		panic("failed to generate url: " + err.Error())
	}

	atcBuild := atc.Build{
		ID:                   build.ID,
		Name:                 build.Name,
		JobName:              build.JobName,
		PipelineID:           build.PipelineID,
		PipelineName:         build.PipelineName,
		PipelineInstanceVars: build.PipelineInstanceVars,
		TeamName:             build.TeamName,
		Status:               string(build.Status),
		APIURL:               apiURL,
	}

	if !build.StartTime.IsZero() {
		atcBuild.StartTime = build.StartTime.Unix()
	}

	if !build.EndTime.IsZero() {
		atcBuild.EndTime = build.EndTime.Unix()
	}

	return atcBuild
}
