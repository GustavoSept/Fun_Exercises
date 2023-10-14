public class CompoundProjector{
    private long initial_Investment = 0;
    private float salary = 0f;
    private float monthlySalaryShare = 0f; // How much the user invests, as a % of the salary
    private float yield = 0f;
    private short investment_years = 0;
    
    private final byte MONTHS_IN_YEARS = 12;


    CompoundProjector(long initial_Investment, float salary, float monthlySalaryShare, float yield, short investment_years){
        this.initial_Investment = initial_Investment;
        this.salary = salary;
        this.monthlySalaryShare = monthlySalaryShare;
        this.yield = yield;
        this.investment_years = investment_years;
    }
    
    public double projectNetWorth(){
        // FORMULA SHORTCUTS
        final double RATE_OVER_MONTHS = this.yield / this.MONTHS_IN_YEARS;
        final double ROM_POW = Math.pow(1+ RATE_OVER_MONTHS, this.MONTHS_IN_YEARS * this.investment_years);
        final float MONTHLY_INVESTMENT = this.salary * this.monthlySalaryShare;



        return (this.initial_Investment * ROM_POW + MONTHLY_INVESTMENT * (ROM_POW-1) / (RATE_OVER_MONTHS));
    }

    public double projectDividends(){
        return projectNetWorth() * this.yield / this.MONTHS_IN_YEARS;
    }

    // Setters
    public void setInitial_Investment(long initial_Investment) {
        this.initial_Investment = initial_Investment;
    }


    public void setSalary(float salary) {
        this.salary = salary;
    }


    public void setMonthlySalaryShare(float monthlySalaryShare) {
        this.monthlySalaryShare = monthlySalaryShare;
    }


    public void setYield(float yield) {
        this.yield = yield;
    }


    public void setInvestment_years(short investment_years) {
        this.investment_years = investment_years;
    }



}